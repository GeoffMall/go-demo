package main

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	buildInfo, _ := debug.ReadBuildInfo()
	version := buildInfo.Main.Version

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"main module version": version,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
