package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/welcome/:nama", func(c *gin.Context) {
		nama := c.Param("nama")
		if nama != "" {
			c.String(http.StatusOK, "Selamat datang "+nama)
		} else {
			c.String(http.StatusOK, "Anonymous")
		}
	})

	// Handling a specific route for "/welcome/"
	router.GET("/welcome/", func(c *gin.Context) {
		c.String(http.StatusOK, "Selamat datang Anonymous")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	router.Run(":" + port)
}
