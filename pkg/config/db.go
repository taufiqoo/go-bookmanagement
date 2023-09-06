package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	d, err := gorm.Open("mysql", "root:@tcp(localhost:3307)/go-book")
	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %s", err))
	}
	db = d
}

func CloseDB() {
	if db != nil {
		err := db.Close()
		if err != nil {
			fmt.Printf("error closing database connection: %s\n", err)
		}
	}
}

func GetDB() *gorm.DB {
	return db
}
