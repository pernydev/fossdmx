package webserver

import (
	"github.com/gin-gonic/gin"
	"github.com/pernydev/fossdmx/pkg/global"
)

func InitRouter() {
	r := gin.New()

	r.GET("/api/ping", Ping)
	r.GET("/api/showfile", GetShowfile)
	r.GET("/api/registry/fixture_types", GetFixtureTypes)
	r.POST("/api/showfile", SetShowfile)

	r.PATCH("/api/patch", MoveFixture)
	r.POST("/api/patch", AddFixture)

	global.Ready <- struct{}{}
	r.Run()
}
