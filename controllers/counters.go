package controllers

import (
	"Url-Counter-Service/models"
	"Url-Counter-Service/services"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
	"strconv"
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
		ctx.Response.SetStatusCode(fasthttp.StatusUnprocessableEntity)
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
	name := string(ctx.QueryArgs().Peek("name"))
	limit, err := strconv.Atoi(string(ctx.QueryArgs().Peek("limit")))

	if err != nil {
		limit = 0
		//log.Print("invalid limit")
		//ctx.Response.SetBody([]byte("invalid limit"))
		//ctx.Response.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		//return
	}
	offset, err := strconv.Atoi(string(ctx.QueryArgs().Peek("offset")))
	if err != nil {
		offset = 0
		//log.Print("invalid offset")
		//ctx.Response.SetBody([]byte("invalid offset"))
		//ctx.Response.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		//return
	}

	counters, err := services.GetCounters(name, limit, offset)

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

	ctx.Response.Header.Set("Content-Type", "json")
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
		return
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
