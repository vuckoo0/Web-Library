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
			log.Println("Missing authorization header")
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

		temp, ok := claims["user_id"].(float64)

		if !ok {
			log.Println("invalid userId claim ", ok, claims["user_id"])
			ctx.JSON(401, gin.H{"error": "invalid userId claim"})
			ctx.Abort()
			return
		}

		userId := int(temp)

		temp, ok = claims["privilege"].(float64)

		if !ok {
			log.Println("invalid privilege claim ", ok, claims["privilege"])
			ctx.JSON(401, gin.H{"error": "invalid privilege claim"})
			ctx.Abort()
			return
		}

		privilege := int(temp)

		ctx.Set("user_id", userId)
		ctx.Set("privilege", privilege)

		ctx.Next()
	}
}

func PrivilegeAuthorization(minPrivilege int) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		privilege := ctx.GetInt("privilege")

		if privilege < 0 {
			log.Println("[PrivilegeAuthorization] Invalid privilege")
			ctx.JSON(403, gin.H{"error": "invalid privilege"})
			ctx.Abort()
			return
		}

		if privilege > minPrivilege {
			log.Println("[PrivilegeAuthorization] Insufficient privilege")
			ctx.JSON(403, gin.H{"error": "insufficient privilege"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
