package services

import (
	"Url-Counter-Service/models"
	"Url-Counter-Service/repositories"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

func CreateCounter(ctx *fasthttp.RequestCtx, counter *models.Counter) error {
	counter.Code = uuid.New().String()

	err := repositories.CreateCounter(counter, context.Background())
	if err != nil {
		return err
	}

	var body []byte
	body, err = json.Marshal(*counter)

	if err != nil {
		return err
	}

	ctx.Response.SetBody(body)
	return err
}
