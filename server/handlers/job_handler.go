package handlers

import "github.com/gin-gonic/gin"

type JobHandler struct {
}

func (h JobHandler) List(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "list",
	})
}
