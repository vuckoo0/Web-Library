package middleware

import (
	"log"
	"main/config"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			log.Println("missing authorization header")
			ctx.JSON(401, gin.H{"error": "missing authorization header"})
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			return []byte(config.Config().JWT_Secret), nil
		})

		if err != nil {
			log.Println(err)
			ctx.JSON(401, gin.H{"error": "invalid or expired token"})
			ctx.Abort()
			return
		}

		if !token.Valid {
			log.Println(err)
			ctx.JSON(401, gin.H{"error": "invalid token"})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Println(err)
			ctx.JSON(401, gin.H{"error": "invalid claims"})
			ctx.Abort()
			return
		}

		userId := claims["user_id"]
		privilege := claims["privilege"]

		ctx.Set("user_id", userId)
		ctx.Set("privilege", privilege)

		ctx.Next()
	}
}
