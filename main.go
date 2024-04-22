package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(cors.Default())

	routerV1 := engine.Group("/api/v1")

	initializeEndpoints(routerV1)

	if err := engine.Run(); err != nil {
		panic(err)
	}
}

func initializeEndpoints(routerV1 *gin.RouterGroup) {

	// Payments
	routerV1.POST("/payments")
	routerV1.GET("/payments/:id")
	routerV1.POST("/payments/:id/capture")

	// Refunds
	routerV1.POST("/refunds")
	routerV1.GET("/refunds/:id")

}
