package controllers

import (
	"Url-Counter-Service/models"
	"Url-Counter-Service/services"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
)

func CreateCounterHandler(ctx *fasthttp.RequestCtx) {
	var counter models.Counter

	err := json.Unmarshal(ctx.PostBody(), &counter)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	err = services.CreateCounter(ctx, &counter)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	log.Print("Added counter")

}

func Redirect(ctx *fasthttp.RequestCtx) {

}
