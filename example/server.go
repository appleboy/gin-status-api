package main

import (
	"github.com/fvbock/endless"
	api "gopkg.in/appleboy/gin-status-api.v1"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/api/status", api.StatusHandler)

	endless.ListenAndServe(":8080", r)
}
