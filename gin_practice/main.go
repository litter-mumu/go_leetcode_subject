package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		_, err := c.Cookie("cookie_key")
		if err != nil {
			c.SetCookie("cookie_key", "It's my cookies", 60,
				"/", "localhost", false, true)
		}
		c.String(200, fmt.Sprintf("Login success"))
	})
	r.Use(func(c *gin.Context) {
		if cookie, err := c.Cookie("cookie_key"); err == nil {
			if cookie == "It's my cookies" {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "权限非法"})
		c.Abort()
		return
	}).GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "home"})
	})
	r.Run(":8000")
}
