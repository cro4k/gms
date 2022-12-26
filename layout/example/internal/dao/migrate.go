package dao

import (
	"github.com/cro4k/gms/layout/example/internal/model"
	"github.com/cro4k/gms/layout/public/db"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var migrations = []*gormigrate.Migration{
	//{
	//	ID: "fix",
	//	Migrate: func(db *gorm.DB) error {
	//		return migration.With(db).
	//			CreateTable("example_table_1").
	//			AddColumn("example_table_2", "example_column").
	//			Error()
	//	},
	//	Rollback: func(db *gorm.DB) error {
	//		return migration.With(db).
	//			DropTable("example_table_1").
	//			DropColumn("example_table_2", "example_column").
	//			Error()
	//	},
	//},
}

func Migrate() {
	m := gormigrate.New(db.DB(), gormigrate.DefaultOptions, migrations)
	m.InitSchema(func(db *gorm.DB) error {
		return db.AutoMigrate(
			&model.User{},
		)
	})
	if err := m.Migrate(); err != nil {
		logrus.Error(err)
	}
}
