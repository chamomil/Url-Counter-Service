package routes

import (
	"Url-Counter-Service/controllers"
	"github.com/fasthttp/router"
)

func CountersRoutes(r *router.Router) {
	r.POST("/counters", controllers.CreateCounterHandler)
	r.GET("/counters/{code}", controllers.Redirect)
	r.GET("/counters", controllers.GetCounters)
	r.GET("/counters/{code}/stats", controllers.RedirectStats)
}
