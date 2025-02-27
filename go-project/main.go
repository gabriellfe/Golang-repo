package main

import (
	"fmt"
	"gabriellfe/config"
	"gabriellfe/rabbit"
	"gabriellfe/routes"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	config.SetupServer()
	//database.SetupDatabase()
	go rabbit.SetupRabbit()

	router := http.NewServeMux()
	routes.Setup(router)

	port := os.Getenv("PORT")
	appName := os.Getenv("APP_NAME")
	slog.Info(fmt.Sprintf("Started %s on port %s", appName, port))
	srv := &http.Server{
		Addr:    port,
		Handler: router,
	}

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
