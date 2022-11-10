package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func applyRoutes(r gin.IRouter) {
	r.GET("hello", func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			name = "World"
		}

		c.String(http.StatusOK, "Hello %s!", name)
	})

	r.GET("hello_json", func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			name = "World"
		}

		c.JSON(http.StatusOK, gin.H{
			"hello": name,
		})
	})
}
