package db

import (
	"github.com/go-errors/errors"
	"github.com/jinzhu/gorm"
	"github.com/pineal-niwan/sensor/logger"
)

//处理db的函数
type HandleDbFunc func(*gorm.DB) error

//封装的事务,没有save point相关,任何一个步骤出错就回滚整个事务
func ExecuteTransaction(db *gorm.DB, anchor string, logger logger.ILogger, handlerList ...HandleDbFunc) (err error) {
	txn := db.Begin()
	err = txn.Error

	if err != nil {
		logger.Errorf(`%s: 开始事务失败: %+v`, anchor, err)
		return err
	}

	defer func() {
		r := recover()
		if r != nil {
			rErr := errors.Wrap(r, 3)
			logger.Errorf(`%s: panic: %+v`, anchor, r)
			logger.Errorf(`%s: panic stack:%+v`, anchor, string(rErr.Stack()))
			err = rErr
		}

		if err != nil {
			logger.Errorf(`%s: 出错,事务需要回滚:%+v`, anchor, err)
			txn.Rollback()
		} else {
			err = txn.Commit().Error
			if err != nil {
				logger.Errorf(`%s: 提交事务出错:%+v`, anchor, err)
			} else {
				logger.Debugf(`%s: 提交事务成功`, anchor)
			}
		}
	}()

	for _, handler := range handlerList {
		err = handler(txn)
		if err != nil {
			return err
		}
	}

	return err
}
