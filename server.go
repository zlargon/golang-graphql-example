package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zlargon/gograph/controller"
	"github.com/zlargon/gograph/handler"
	"github.com/zlargon/gograph/middleware"
	"github.com/zlargon/gograph/service"
)

var (
	jwtService    = service.NewJwtService()
	jwtController = controller.NewJwtController(jwtService)
)

func main() {
	// setup gin
	server := gin.Default()
	server.GET("/jwt", middleware.BasicAuth(), jwtController.GenerateToken)
	server.GET("/", handler.PlaygroundHandler())
	server.POST("/query", handler.GraphqlHandler())

	port := getPort()
	server.Run(":" + port)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
}

func getPort() string {
	const defaultPort = "8080"
	port := os.Getenv("PORT")
	if port == "" {
		return defaultPort
	}
	return port
}
