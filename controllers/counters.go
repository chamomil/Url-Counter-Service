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
		log.Print(err.Error())
		ctx.Response.SetBody([]byte("error in params conversion"))
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	err = services.CreateCounter(&counter)
	if err != nil {
		log.Print(err.Error())
		ctx.Response.SetBody([]byte("error in creating counter"))
		ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	var body []byte
	body, err = json.Marshal(counter)

	if err != nil {
		log.Print(err.Error())
		ctx.Response.SetBody([]byte("error in response body conversion"))
		ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.Response.SetBody(body)

	log.Print("Added counter")

}

func Redirect(ctx *fasthttp.RequestCtx) {
	code, ok := ctx.UserValue("code").(string)

	if !ok {
		log.Print("invalid code")
		ctx.Response.SetBody([]byte("invalid code"))
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	url, err := services.GetUrlByCode(code)
	if err != nil {
		log.Print(err.Error())
		ctx.Response.SetBody([]byte("error getting url from db"))
		ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.Redirect(url, fasthttp.StatusPermanentRedirect)
	log.Printf("Redirect to %s", url)
}

func GetCounters(ctx *fasthttp.RequestCtx) {
	var name string

	_ = json.Unmarshal(ctx.Request.Body(), &name)

	counters, err := services.GetCounters(name)

	if err != nil {
		log.Print(err.Error())
		ctx.Response.SetBody([]byte(err.Error()))
		ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	var body []byte
	body, err = json.Marshal(*counters)

	if err != nil {
		log.Print(err.Error())
		ctx.Response.SetBody([]byte("error in response body conversion"))
		ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.Response.SetBody(body)

}
