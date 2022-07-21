package api

import "github.com/gin-gonic/gin"

func error(message string) gin.H {
	return gin.H{
		"error": message,
	}
}
