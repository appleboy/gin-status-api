package status

import (
	"net/http"

	api "github.com/fukata/golang-stats-api-handler"
	"github.com/gin-gonic/gin"
)

// GinHandler is gin handle for get system status.
func GinHandler(c *gin.Context) {
	c.JSON(http.StatusOK, api.GetStats())
}
