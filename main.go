package main

import (
	"Url-Counter-Service/config"
	"Url-Counter-Service/db"
	"Url-Counter-Service/middlewares"
	"Url-Counter-Service/routes"
	"context"
	"fmt"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
)

func Initialize() {
	configData, err := config.ReadConfig("config.yml")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.InitConnection(context.Background(), &configData.Postgres)
	if err != nil {
		log.Fatal(err.Error())
	}
	if configData.Postgres.RunMigrations {
		err = db.RunMigrations("migrations")
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

func main() {
	Initialize()
	r := router.New()
	routes.CountersRoutes(r)

	port := config.Data.ServerPort
	address := fmt.Sprintf(":%d", port)
	log.Printf("Server running on port %d. Try http://localhost%s/counters", port, address)
	log.Fatal(fasthttp.ListenAndServe(address, middlewares.Apply(r.Handler, middlewares.Auth, middlewares.Logger)))
}
