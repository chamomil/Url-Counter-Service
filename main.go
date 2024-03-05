package main

import (
	"Url-Counter-Service/config"
	"Url-Counter-Service/db"
	"Url-Counter-Service/routes"
	"context"
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

func Finalize() {
	if err := db.CloseConnection(context.Background()); err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	Initialize()
	defer Finalize()
	r := router.New()
	routes.CountersRoutes(r)

	log.Print("Server running on port 8080. Try http://localhost:8080/counters")
	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}
