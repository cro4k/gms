package db

import (
	"github.com/cro4k/gms/layout/public/global"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(
		mysql.New(mysql.Config{DSN: global.C().DB.DSN()}),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{SingularTable: true},
		},
	)
	if err != nil {
		logrus.Fatal(err)
	}
}

func DB() *gorm.DB {
	return db
}
