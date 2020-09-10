package jwt

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		bearerToken := header[len("Bearer"):]
		token, err := NewJWTService().Validate(bearerToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized, Please login and try again",
			})
			return
		}

		if token.Valid {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized, Please login and try again",
			})
			return
		}
	}
}
