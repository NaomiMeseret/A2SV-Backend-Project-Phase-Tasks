package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)
var JwtSecret = []byte("change")

func AuthRequired() gin.HandlerFunc{
	return func(c *gin.Context){
		header :=c.GetHeader("Authorization")
		parts := strings.Split(header, " ")
		if len(parts)!=2 || strings.ToLower(parts[0])!="bearer"{
			c.IndentedJSON(http.StatusUnauthorized , gin.H{"error":"Auth required"})
			c.Abort()
			return
		}
	token , err := jwt.Parse(parts[1],func(t *jwt.Token)(interface{} , error){
		if _, ok := t.Method.(*jwt.SigningMethodHMAC);!ok{
			return nil , fmt.Errorf("unexpected signing method")
		}
		return JwtSecret , nil
	})
	if err != nil || !token.Valid{
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error":"Invaild token"})
		c.Abort()
		return
	}
	claims , ok := token.Claims.(jwt.MapClaims)
	if !ok{
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error" :"Invaild token"})
		c.Abort()
		return
	}
	c.Set("userID" , claims["userID"])
	c.Set("role" , claims["role"])
	c.Next()
}
}
func AdminOnly () gin.HandlerFunc{
	return func(c *gin.Context){
		role , _ := c.Get("role")
		if role != "admin"{
			c.IndentedJSON(http.StatusForbidden , gin.H{"error":"Admins only"})
			c.Abort()
			return
		}
		c.Next()
	}
}