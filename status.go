package status

import (
	api "github.com/fukata/golang-stats-api-handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

// StatusHandler is gin handle for get system status.
func StatusHandler(c *gin.Context) {
	c.JSON(http.StatusOK, api.GetStats())
}
