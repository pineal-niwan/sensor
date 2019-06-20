package db

import (
	"github.com/jinzhu/gorm"
	"github.com/pineal-niwan/sensor/logger"
)

type HandleDbFunc func(*gorm.DB) error

func ExecuteTransaction(db *gorm.DB, handler HandleDbFunc, anchor string, logger logger.ILogger) error {
	txn := db.Begin()
	err := txn.Error
	if err != nil {
		logger.Errorf(`%s: 开始事务失败: %+v`, anchor, err)
		return err
	}
	defer func() {
		err = recover()
	}()

}