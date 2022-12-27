package controller

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zlargon/gograph/service"
)

// ==============================================
// Interface
// ==============================================
type JwtController interface {
	GenerateToken(ctx *gin.Context)
}

// ==============================================
// Implementation
// ==============================================
type controller struct {
	service service.JwtService
}

func NewJwtController(service service.JwtService) JwtController {
	return &controller{
		service: service,
	}
}

func (c *controller) GenerateToken(ctx *gin.Context) {
	// Get username from Basic authorization header
	const BASIC_SCHEMA = "Basic "
	authHeader := ctx.GetHeader("authorization")
	tokenString := authHeader[len(BASIC_SCHEMA):]
	usernameAndPassword, _ := base64.URLEncoding.DecodeString(tokenString)
	username := strings.Split(string(usernameAndPassword), ":")[0]

	// Generate JWT token by username
	token, err := c.service.GenerateToken(username)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
