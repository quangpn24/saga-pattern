package main

import (
	"fmt"
	"log"
	"payment-service/config"

	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("cannot load config: %v\n", err)
	}

	//mysqlDsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true", cfg.DB.User, cfg.DB.Pass, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable", cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.DBName, cfg.DB.Port)

	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		log.Fatalf("cannot connecting to db: %v\n", err)
	}

	dbConn = dbConn.Debug()

	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	sqlDB, err := dbConn.DB()
	if err != nil {
		log.Fatalf("cannot get db connection: %v\n", err)
	}

	total, err := migrate.Exec(sqlDB, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalf("cannot execute migration: %v\n", err)
	}

	log.Printf("applied %d migrations\n", total)
}
