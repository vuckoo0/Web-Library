package middleware

import (
	"log"
	"main/config"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			log.Println(authHeader)
			c.JSON(401, gin.H{"error": "missing authorization header"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			return []byte(config.Config().JWT_Secret), nil
		})

		if err != nil {
			log.Println(err)
			c.JSON(401, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		if !token.Valid {
			log.Println(err)
			c.JSON(401, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Println(err)
			c.JSON(401, gin.H{"error": "invalid claims"})
			c.Abort()
			return
		}

		userId := claims["user_id"]
		privilege := claims["privilege"]

		c.Set("user_id", userId)
		c.Set("privilege", privilege)

		c.Next()
	}
}
