package middleware

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/zlargon/gograph/service"
)

var (
	jwtService = service.NewJwtService()
)

func AuthorizeJwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		const BEARER_SCHEMA = "Bearer "
		authHeader := ctx.GetHeader("authorization")
		if headerLen := len(authHeader); headerLen == 0 || headerLen < len(BEARER_SCHEMA) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing Authorization header",
			})
			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := jwtService.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token is invalid",
			})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		log.Println(claims)
	}
}
