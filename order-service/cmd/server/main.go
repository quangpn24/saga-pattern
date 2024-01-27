package main

import (
	"fmt"
	"log"
	"net"
	"order-service/config"
	"order-service/handler/consumer"
	serviceHttp "order-service/handler/http"
	"order-service/repository"
	"order-service/usecase"

	"github.com/sirupsen/logrus"
	"github.com/soheilhy/cmux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const VERSION = "1.0.0"

// @title Example API
// @version 1.0

// @BasePath /api
// @schemes http http

// @securityDefinitions.apikey AuthToken
// @in header
// @name Authorization

// @description Transaction API.
func main() {
	cfg, _ := config.LoadConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable", cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.DBName, cfg.DB.Port)

	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		log.Fatalf("cannot connecting to db: %v\n", err)
	}

	dbConn = dbConn.Debug()

	repo := repository.New(dbConn)
	uc := usecase.New(repo)

	go runConsumer(uc, cfg)
	executeServer(uc, cfg)
}
func runConsumer(uc *usecase.UseCase, cfg *config.Config) {
	// to consume messages
	fmt.Println("Start consumer")
	con := consumer.NewConsumer(uc, cfg)
	con.Consume()
}
func executeServer(useCase *usecase.UseCase, cfg *config.Config) {
	fmt.Println(`cfg.Port: `, cfg.Port)
	l, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		logrus.Fatal(err)
	}

	m := cmux.New(l)
	httpL := m.Match(cmux.HTTP1Fast())
	errs := make(chan error)

	// http
	{
		h := serviceHttp.NewHTTPHandler(useCase, cfg)
		go func() {
			h.Listener = httpL
			errs <- h.Start("")
		}()
	}

	go func() {
		errs <- m.Serve()
	}()

	err = <-errs
	if err != nil {
		logrus.Fatal(err)
	}
}
