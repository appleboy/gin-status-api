package main

import (
	api "github.com/appleboy/gin-status-api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/api/status", api.GinHandler)

	r.Run() // listen and serve on 0.0.0.0:8080
}
