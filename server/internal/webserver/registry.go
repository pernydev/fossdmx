package webserver

import (
	"github.com/gin-gonic/gin"
	"github.com/pernydev/fossdmx/pkg/registry"
)

func GetFixtureTypes(c *gin.Context) {
	c.JSON(200, registry.FixtureTypes)
}
