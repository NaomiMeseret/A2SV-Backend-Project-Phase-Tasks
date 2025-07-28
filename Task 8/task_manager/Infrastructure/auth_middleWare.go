package infrastructure

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//AuthRequired is a Gin middleware
func AuthRequired() gin.HandlerFunc{
	return func(c *gin.Context) {
        header := c.GetHeader("Authorization")
        parts := strings.Split(header, " ")
        if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
            c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Auth required"})
            c.Abort()
            return
        }
        token, err := jwt.Parse(parts[1], func(t *jwt.Token) (interface{}, error) {
            if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method")
            }
			return jwtSecret , nil
		})
		if err != nil || !token.Valid{
			c.IndentedJSON(http.StatusUnauthorized ,gin.H{"error":"Invalid token"})
			c.Abort()
			return 
		}
		claims , ok := token.Claims.(jwt.MapClaims)
		if !ok{
			c.IndentedJSON(http.StatusUnauthorized , gin.H{"error":"Invalid token"})
			c.Abort()
			return
		}
		c.Set("user_id" , claims["user_id"])
		c.Set("role" , claims["role"])
		c.Next()
    }
}