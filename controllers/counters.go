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
		ctx.Response.SetBody([]byte(err.Error()))
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	err = services.CreateCounter(&counter)
	if err != nil {
		log.Print(err.Error())
		ctx.Response.SetBody([]byte(err.Error()))
		ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	var body []byte
	body, err = json.Marshal(counter)

	if err != nil {
		log.Print(err.Error())
		ctx.Response.SetBody([]byte(err.Error()))
		ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.Response.SetBody(body)

	ctx.Response.SetStatusCode(fasthttp.StatusCreated)

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
		ctx.Response.SetBody([]byte(err.Error()))
		ctx.Response.SetStatusCode(fasthttp.StatusNotFound)
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
		ctx.Response.SetBody([]byte(err.Error()))
		ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.Response.SetBody(body)
	ctx.Response.SetStatusCode(fasthttp.StatusFound)
}

func RedirectStats(ctx *fasthttp.RequestCtx) {
	code, ok := ctx.UserValue("code").(string)

	if !ok {
		log.Print("invalid code")
		ctx.Response.SetBody([]byte("invalid code"))
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	count, err := services.GetRedirects(code)
	if err != nil {
		log.Print(err.Error())
		ctx.Response.SetBody([]byte(err.Error()))
		ctx.Response.SetStatusCode(fasthttp.StatusNotFound)
	}

	var body []byte
	body, err = json.Marshal(count)
	if err != nil {
		log.Print(err.Error())
		ctx.Response.SetBody([]byte(err.Error()))
		ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	ctx.Response.SetBody(body)
	ctx.Response.SetStatusCode(fasthttp.StatusFound)
}
