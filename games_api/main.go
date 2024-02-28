package main

import (
	"games-api/api"
	"games-api/data"
	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func newServer(gameApi *api.API) *gin.Engine {
	swagger, err := api.GetSwagger()
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.Use(middleware.OapiRequestValidator(swagger))
	api.RegisterHandlers(router, gameApi)
	return router
}

func main() {
	err := data.NewDB()
	if err != nil {
		panic(err)
	}
	server := newServer(api.NewAPI())
	server.Run(":8080")
}
