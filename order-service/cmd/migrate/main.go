package main

import (
	"log"

	"github.com/Finatext/rakugan-cms-server/pkg/applogger"
	"github.com/Finatext/rakugan-cms-server/pkg/config"
	"github.com/Finatext/rakugan-cms-server/pkg/mysql"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	logger, err := applogger.NewAppLogger()
	if err != nil {
		log.Fatalf("cannot init logger: %v\n", err)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatalf("cannot load config: %v\n", err)
	}

	dbConn, err := mysql.Open(cfg)
	if err != nil {
		logger.Fatalf("cannot connecting to db: %v\n", err)
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	sqlDB, err := dbConn.DB()
	if err != nil {
		logger.Fatalf("cannot get db connection: %v\n", err)
	}

	total, err := migrate.Exec(sqlDB, "mysql", migrations, migrate.Up)
	if err != nil {
		logger.Fatalf("cannot execute migration: %v\n", err)
	}

	logger.Infof("applied %d migrations\n", total)
}
