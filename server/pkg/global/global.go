package global

import (
	"os"
	"time"

	"github.com/adrg/xdg"
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	dmx "github.com/pernydev/fossdmx/pkg/dmx"
	"github.com/pernydev/fossdmx/pkg/models"
)

var (
	DMX            = dmx.NewDMX()
	ShowfileLoaded = false
	Showfile       models.Showfile
	Gin            *gin.Engine
	Log            = log.NewWithOptions(os.Stderr, log.Options{
		TimeFormat:   time.Kitchen,
		ReportCaller: true,
	})
	DataFolder = xdg.DataHome + "/fossdmx/"
	Ready      = make(chan struct{})
)
