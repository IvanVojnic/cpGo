package main

import (
	"github.com/IvanVojnic/cpGo.git"
	"github.com/IvanVojnic/cpGo.git/pkg/handler"
	"github.com/IvanVojnic/cpGo.git/pkg/repository"
	"github.com/IvanVojnic/cpGo.git/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(cpGo.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("configs")
	return viper.ReadInConfig()
}
