package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Database() *gorm.DB {
	db, err := gorm.Open("mysql", "root:dev@/shortener?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Failed to connect to the database")
	}

	return db
}
