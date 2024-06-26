// @/config/database.go
package config

import (
	"fmt"
	"ncip48/go-fiber-ecommerce/entities"
	"ncip48/go-fiber-ecommerce/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var DATABASE_URI string = fmt.Sprintf(
	"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	utils.GoDotEnvVariable("DB_USER"),
	utils.GoDotEnvVariable("DB_PASSWORD"),
	utils.GoDotEnvVariable("DB_HOST"),
	utils.GoDotEnvVariable("DB_PORT"),
	utils.GoDotEnvVariable("DB_NAME"),
)

func Connect() error {
	var err error

	Database, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&entities.User{})

	return nil
}
