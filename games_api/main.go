package main

import (
	"games-api/api"
	"games-api/data"
	"github.com/gin-gonic/gin"
	middleware "github.com/oapi-codegen/gin-middleware"
	ginmiddleware "github.com/zsais/go-gin-prometheus"
)

func newServer(gameApi *api.API) *gin.Engine {
	swagger, err := api.GetSwagger()
	if err != nil {
		panic(err)
	}
	router := gin.Default()
	prom := ginmiddleware.NewPrometheus("gin")
	prom.Use(router)
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
