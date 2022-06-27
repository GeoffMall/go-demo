package main

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	buildInfo, _ := debug.ReadBuildInfo()

	mainModuleVersion := buildInfo.Main.Version
	vscModified, vcsRevision, vcsTime := "", "", ""

	for _, setting := range buildInfo.Settings {
		if setting.Key == "vcs.revision" {
			vcsRevision = setting.Value
		}
		if setting.Key == "vcs.time" {
			vcsTime = setting.Value
		}
		if setting.Key == "vcs.modified" {
			vscModified = setting.Value
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"version":     mainModuleVersion,
			"vcsModified": vscModified,
			"vcsRevision": vcsRevision,
			"vcsTime":     vcsTime,
		})
	})
	_ = r.Run()
}
