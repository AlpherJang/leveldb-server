package handlers

import (
	"github.com/gin-gonic/gin"
	"leveldb-server/configs"
	"net/http"
)

type ReaderHandler struct {
}
type ReadOneRes struct {
	Key string `uri:"key" binding:"required"`
}

func (h ReaderHandler) ReadOne(c *gin.Context) {
	var res ReadOneRes
	if err := c.ShouldBindUri(&res); err != nil {
		c.JSON(http.StatusBadRequest, configs.BuildResult(5000, err.Error(), nil))
	}
	if len(res.Key) == 0 {
		c.JSON(http.StatusBadRequest, configs.BuildResult(10000, "key and value must be non-empty", nil))
		return
	}
	data, err := configs.Db.Get(res.Key)
	if err != nil {
		c.JSON(http.StatusBadRequest, configs.BuildResult(30000, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, configs.BuildResult(20000, "Success", data))
}
