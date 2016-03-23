package status

import (
	"github.com/gin-gonic/gin"
	api "github.com/fukata/golang-stats-api-handler"
)

func StatusHandler(c *gin.Context) {
	c.JSON(200, api.GetStats())
}
