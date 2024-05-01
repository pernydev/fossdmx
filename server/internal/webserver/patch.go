package webserver

import (
	"github.com/gin-gonic/gin"
	"github.com/pernydev/fossdmx/pkg/global"
	"github.com/pernydev/fossdmx/pkg/models"
)

func MoveFixture(c *gin.Context) {
	var data struct {
		FixtureID string         `json:"fixture_id"`
		Channel   models.Channel `json:"channel"`
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	fixture := global.Showfile.Fixtures[data.FixtureID]
	if fixture == nil {
		c.JSON(404, gin.H{"error": "Fixture not found"})
		return
	}

	fixture.Channel = data.Channel
	c.Status(204)
}

func AddFixture(c *gin.Context) {
	var data struct {
		Fixture models.Fixture `json:"fixture"`
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	global.Showfile.Fixtures[data.Fixture.ID] = &data.Fixture
	c.JSON(200, data.Fixture)
}
