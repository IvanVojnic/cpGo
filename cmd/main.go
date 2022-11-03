package main

import (
	"github.com/IvanVojnic/cpGo.git"
	"github.com/IvanVojnic/cpGo.git/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(cpGo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error", err.Error())
	}
}
