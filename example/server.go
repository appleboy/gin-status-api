package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	api "github.com/appleboy/gin-status-api"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/api/status", api.StatusHandler)

	endless.ListenAndServe(":8080", r)
}
