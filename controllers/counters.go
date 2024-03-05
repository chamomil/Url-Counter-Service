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

	err = services.CreateCounter(&counter)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	var body []byte
	body, err = json.Marshal(counter)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	ctx.Response.SetBody(body)

	log.Print("Added counter")

}

func Redirect(ctx *fasthttp.RequestCtx) {
	code, ok := ctx.UserValue("code").(string)

	if !ok {
		log.Fatal("invalid code")
		return
	}

	url, err := services.GetUrlByCode(code)
	if err != nil {
		log.Fatal(err.Error())
	}

	ctx.Redirect(url, fasthttp.StatusPermanentRedirect)
	log.Printf("Redirect to %s", url)
}
