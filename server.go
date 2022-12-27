package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zlargon/gograph/handler"
	"github.com/zlargon/gograph/middleware"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// setup gin
	server := gin.Default()
	server.Use(middleware.BasicAuth())
	server.GET("/", handler.PlaygroundHandler())
	server.POST("/query", handler.GraphqlHandler())
	server.Run(":" + port)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
}
