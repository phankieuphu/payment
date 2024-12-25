package databases

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbInstance *gorm.DB
	dbError    error
)

func GetWriteInstance() *gorm.DB {
	if dbInstance == nil {
		InitDB()
	}
	return dbInstance

}

func InitDB() {
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dbInstance, dbError = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dbError != nil {
		panic("Failed to connect to database")
	}
}
