package webserver

import (
	"github.com/gin-gonic/gin"
	"github.com/pernydev/fossdmx/pkg/global"
	"github.com/pernydev/fossdmx/pkg/showfile"
)

func GetShowfile(c *gin.Context) {
	if !global.ShowfileLoaded {
		c.JSON(404, gin.H{"error": "No showfile loaded"})
		return
	}
	c.JSON(200, global.Showfile)
}

func SetShowfile(c *gin.Context) {
	var data struct {
		Filename string `json:"filename"`
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := showfile.LoadFromFile(data.Filename)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.Status(204)
}
