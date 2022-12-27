package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zlargon/gograph/http"
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
	server.GET("/", http.PlaygroundHandler())
	server.POST("/query", http.GraphqlHandler())
	server.Run(":" + port)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
}
