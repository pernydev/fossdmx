package webserver

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.New()

	r.GET("/api/ping", Ping)
	r.GET("/api/showfile", GetShowfile)
	r.GET("/api/registry/fixture_types", GetFixtureTypes)
	r.POST("/api/showfile", SetShowfile)

	r.POST("/api/patch", Patch)

	r.Run()
}
