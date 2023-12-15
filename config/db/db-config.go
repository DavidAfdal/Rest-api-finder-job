package dbconfig

import (
	"fmt"

	usermodel "finder-job/model/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host = "localhost"
	port = 5432
	user = "davidafdal"
	password = "gelang123"
	dbName = "test"
)

func ConnectDB() *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&usermodel.User{})
	return db
}