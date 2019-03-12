package config

import (
	_ "github.com/go-sql-driver/mysql" // Package to connect with MySQL or MariaDB
	"github.com/jinzhu/gorm"           // Main GORM package
)

// Database it return db instance from Gorm
func Database() *gorm.DB {
	db, err := gorm.Open("mysql", "root:dev@/shortener?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Failed to connect to the database")
	}

	return db
}
