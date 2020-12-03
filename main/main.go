package main

import (
	"github.com/shijuvar/gokit/examples/training-api/controller"
	"github.com/shijuvar/gokit/examples/training-api/mapstore"
	"github.com/shijuvar/gokit/examples/training-api/router"
	"go.uber.org/zap"
	"log"
	"net/http"
)


func main() {
	logger, _ := zap.NewProduction() // Create Uber's Zap logger
	controller := &controller.CustomerController{
		Store:  mapstore.NewMapStore(),
		Logger: logger,
	}
	r := router.InitializeRoutes(controller)
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe() // Run the http server
}