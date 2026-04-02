package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	transportAuth "github.com/superdima3000/transport-auth"
	"github.com/superdima3000/transport-auth/db"
	"github.com/superdima3000/transport-auth/pkg/controller"
	"github.com/superdima3000/transport-auth/pkg/repository"
	"github.com/superdima3000/transport-auth/pkg/service"
)

// @title MIFARE API
// @version 0.1
// @description RESTful API server for MIFARE transactions authorization

// @host      localhost:8888
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Введите токен в формате: "Bearer <token>"

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if err := initConfig(); err != nil {
		logrus.Fatal(err.Error())
	}
	database := db.New(viper.GetString("db.path"))
	defer database.Close()

	srv := new(transportAuth.Server)
	repos := repository.NewRepository(database)
	services := service.NewService(repos)
	handler := controller.NewHandler(services)

	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatal(err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configuration")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
