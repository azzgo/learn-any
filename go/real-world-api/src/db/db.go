package db

import (
	"real-world-api/src/common"

	"github.com/jinzhu/gorm"
)


func openDB() (*gorm.DB, error) {
	var USERNAME = common.Config["MYSQL_USERNAME"]
	var PASSWORD = common.Config["MYSQL_PASSWORD"]
	return gorm.Open("mysql", USERNAME+":"+PASSWORD+"@(localhost:3306)/real-world-db?&parseTime=True")
}

// GetDB godoc
func GetDB() *gorm.DB {
	db, err := openDB()

	if err != nil {
		panic("connect Db failed: " + err.Error())
	}

	return db
}
