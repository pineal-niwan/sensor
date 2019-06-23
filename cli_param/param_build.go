package cli_param

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pineal-niwan/sensor/cypher"
	"gopkg.in/urfave/cli.v1"
	"time"
)

const (
	DbLinkFmt = `%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local`
)

//hash cycherKey
func hashCypherKey(cypherKey string) (string, error) {
	decryptKey, err := cypher.Sha1String(`yc#h*zzx-bbd` + cypherKey + `axk,mzl9901kxk1+_[]1\z`)
	if err != nil {
		return decryptKey, err
	}
	decryptKey, err = cypher.Sha256String(`BVK-IOK-8192-OL-pqw` + cypherKey + `YJp)BU_-O-KDW-OQ`)
	if err != nil {
		return decryptKey, err
	}
	decryptKey, err = cypher.Sha512String(`XVP,.zrt+qa=zd9f34qw` + cypherKey + `/xvk-.vsw-`)
	if err != nil {
		return decryptKey, err
	}
	decryptKey, err = cypher.Md5String(decryptKey)
	return decryptKey, err
}

//建立db
func BuildDb(c *cli.Context) (*gorm.DB, error) {
	cypherKey := GetCypherKey(c)
	decryptKey, err := hashCypherKey(cypherKey)
	if err != nil {
		return nil, err
	}

	dbPass := GetDbPass(c)
	dbPass, err = cypher.Decrypt([]byte(decryptKey), dbPass)
	if err != nil {
		return nil, err
	}

	dbLink := fmt.Sprintf(DbLinkFmt, GetDbUser(c), dbPass, GetDbAddress(c), GetDbSchema(c))
	db, err := gorm.Open(`mysql`, dbLink)
	if err != nil {
		return db, err
	}
	db.DB().SetMaxOpenConns(GetDbMaxPoolSize(c))
	db.DB().SetMaxIdleConns(GetDbMaxPoolSize(c))
	dbMaxLifeTime := time.Duration(GetDbMaxLifeTime(c)) * time.Second
	db.DB().SetConnMaxLifetime(dbMaxLifeTime)

	return db, err
}
