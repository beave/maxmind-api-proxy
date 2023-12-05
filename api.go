package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	)

func Authenticate_API() gin.HandlerFunc {
	return func(c *gin.Context) {

		api_key := c.GetHeader("API_KEY")
		
		if api_key == "" { 
			c.JSON(http.StatusOK, gin.H{"error": "api authentication failed"})
			c.Abort()
			return
		}

		if api_key != Config.API_Key {
			c.JSON(http.StatusOK, gin.H{"error": "api authentication failed"})
			c.Abort()
			return
		}
	}
}
