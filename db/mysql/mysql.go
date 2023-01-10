package mysql

// REF: https://gorm.io/zh_CN/docs/connecting_to_the_database.html

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Open(dsn string, opts ...gorm.Option) (db *gorm.DB, err error) {
	return gorm.Open(mysql.Open(dsn), opts...)
}

func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.Close()
	return nil
}
