package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"projectIntern/pkg/config"
)

func DBInit(env *config.Env) (*gorm.DB, error) {
	username := env.DBUsername
	password := env.DBPassword
	host := env.DBHost
	port := env.DBPort
	database := env.DBName

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	return db, err
}
