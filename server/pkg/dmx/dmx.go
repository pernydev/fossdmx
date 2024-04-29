package dmx

import (
	"github.com/pernydev/fossdmx/pkg/models"
)

type DMX struct {
	Channels map[models.Channel]uint8 `json:"channels"`
}

func (d *DMX) Update() {
	// TODO: broadcast update further
}

func NewDMX() *DMX {
	return &DMX{
		Channels: make(map[models.Channel]uint8),
	}
}
