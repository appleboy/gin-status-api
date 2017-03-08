package status

import (
	"gopkg.in/gin-gonic/gin.v1"
	api "gopkg.in/fukata/golang-stats-api-handler.v1"
	"net/http"
)

// StatusHandler is gin handle for get system status.
func StatusHandler(c *gin.Context) {
	c.JSON(http.StatusOK, api.GetStats())
}
