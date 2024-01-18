package main

import (
	"delivery-service/config"
	"fmt"
	"log"

	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("cannot load config: %v\n", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true", cfg.DB.User, cfg.DB.Pass, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true", confMysql.DBUser, confMysql.DBPass, confMysql.DBHost, confMysql.DBPort)

	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		log.Fatalf("cannot connecting to db: %v\n", err)
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	sqlDB, err := dbConn.DB()
	if err != nil {
		log.Fatalf("cannot get db connection: %v\n", err)
	}

	total, err := migrate.Exec(sqlDB, "mysql", migrations, migrate.Up)
	if err != nil {
		log.Fatalf("cannot execute migration: %v\n", err)
	}

	log.Printf("applied %d migrations\n", total)
}
