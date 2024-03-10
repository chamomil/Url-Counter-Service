package middlewares

import (
	"Url-Counter-Service/services"
	"encoding/base64"
	"fmt"
	"github.com/valyala/fasthttp"
	"strings"
)

func Auth(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		fmt.Println("auth")
		authHeader := string(ctx.Request.Header.Peek("Authorization"))
		if authHeader == "" {
			ctx.SetStatusCode(fasthttp.StatusUnauthorized)
			return
		}
		const prefix = "Basic "
		byteCredentials, err := base64.StdEncoding.DecodeString(authHeader[len(prefix):])
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			return
		}
		username, password, ok := strings.Cut(string(byteCredentials), ":")
		if !ok {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			return
		}
		if services.IsCredentialsValid(username, password) {
			handler(ctx)
			return
		}
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
	}
}
