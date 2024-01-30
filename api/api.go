package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"async_queue/utils"
)

func StartApi(config *utils.Configuration) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(fmt.Sprintf("%s:%s", config.ApiHost, config.ApiPort))
}
