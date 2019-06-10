package main

import (
	api "github.com/appleboy/gin-status-api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/api/status", api.StatusHandler)

	r.Run() // listen and serve on 0.0.0.0:8080
}
