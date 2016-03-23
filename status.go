package status

import (
	api "github.com/fukata/golang-stats-api-handler"
	"github.com/gin-gonic/gin"
)

func StatusHandler(c *gin.Context) {
	c.JSON(200, api.GetStats())
}
