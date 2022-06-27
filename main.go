package main

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	buildInfo, _ := debug.ReadBuildInfo()

	vcsRevision, vcsTime := "", ""

	for _, setting := range buildInfo.Settings {
		if setting.Key == "vcs.revision" {
			vcsRevision = setting.Value
		}
		if setting.Key == "vcs.time" {
			vcsTime = setting.Value
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"vcsRevision": vcsRevision,
			"vcsTime":     vcsTime,
		})
	})
	_ = r.Run()
}
