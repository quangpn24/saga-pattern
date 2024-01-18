package main

import (
	"fmt"
	"log"
	"net"
	"order-service/config"
	"order-service/repository"
	"order-service/usecase"

	serviceHttp "order-service/handler/http"

	"github.com/sirupsen/logrus"
	"github.com/soheilhy/cmux"
	"gorm.io/driver/mysql"
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
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true", cfg.DB.User, cfg.DB.Pass, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true", confMysql.DBUser, confMysql.DBPass, confMysql.DBHost, confMysql.DBPort)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.New(db)
	uc := usecase.New(repo)

	//go func() {
	//	h.Listener = httpL
	//	errs <- h.Start("")
	//}()
	executeServer(uc)
}
func executeServer(useCase *usecase.UseCase) {
	cfg, _ := config.LoadConfig()

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
		h := serviceHttp.NewHTTPHandler(useCase)
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
