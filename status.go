package status

import (
	"net/http"

	"github.com/gin-gonic/gin"
	api "gopkg.in/fukata/golang-stats-api-handler.v1"
)

// StatusHandler is gin handle for get system status.
func StatusHandler(c *gin.Context) {
	c.JSON(http.StatusOK, api.GetStats())
}
