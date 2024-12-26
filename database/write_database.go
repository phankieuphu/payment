package databases

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbInstance *gorm.DB
	dbError    error
)

func GetWriteInstance() *gorm.DB {
	if dbInstance == nil {
		InitWriteDB()
	}
	return dbInstance

}

func InitWriteDB() {
	dsn := os.Getenv("DATABASE_WRITE_URL")
	// dsn := "readuser:readpassword@tcp(db-read:3306)/payment"
	fmt.Println(dsn)
	dbInstance, dbError = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if dbError != nil {
		panic("Failed to connect to database")
	}
	log.Print("Write database connected")
}
