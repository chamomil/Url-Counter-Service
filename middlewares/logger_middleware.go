package middlewares

import (
	"github.com/valyala/fasthttp"
	"log"
	"os"
	"time"
)

var output = log.New(os.Stdout, "", log.LstdFlags)

func Logger(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		begin := time.Now()
		handler(ctx)
		end := time.Now()
		total := end.Sub(begin)

		output.Printf("%v | %s %s - %v - %v \n Request body: %s \n Request headers: %s \n Response body: %s \n Response headers: %s",
			ctx.RemoteAddr(),
			ctx.Method(),
			ctx.RequestURI(),
			ctx.Response.Header.StatusCode(),
			total,
			string(ctx.Request.Body()),
			ctx.Request.Header.String(),
			string(ctx.Response.Body()),
			ctx.Response.Header.String(),
		)
	}
}
