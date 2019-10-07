package handlers

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"leveldb-server/configs"
	"net/http"
)

type WriterHandler struct {
}

type WriteItem struct {
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
}

func (h WriterHandler) PutOne(c *gin.Context) {
	var writeItem WriteItem
	if err := c.ShouldBindJSON(&writeItem); err != nil {
		c.JSON(http.StatusBadRequest, configs.BuildResult(5000, err.Error(), nil))
		return
	}
	if len(writeItem.Key) == 0 || len(writeItem.Value) == 0 {
		c.JSON(http.StatusBadRequest, configs.BuildResult(10000, "key and value must be non-empty", nil))
		return
	}
	value, err := base64.StdEncoding.DecodeString(writeItem.Value)
	if err != nil {
		c.JSON(http.StatusBadRequest, configs.BuildResult(40000, err.Error(), nil))
		return
	}
	err = configs.Db.PutOne(writeItem.Key, value)
	if err != nil {
		c.JSON(http.StatusBadRequest, configs.BuildResult(30000, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, configs.BuildResult(20000, "Success", nil))
}
