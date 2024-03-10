package middlewares

import (
	"github.com/valyala/fasthttp"
)

type Middleware func(handler fasthttp.RequestHandler) fasthttp.RequestHandler

func Apply(handler fasthttp.RequestHandler, middlewares ...Middleware) fasthttp.RequestHandler {
	resultHandler := handler
	for _, middleware := range middlewares {
		resultHandler = middleware(resultHandler)
	}
	return resultHandler
}
