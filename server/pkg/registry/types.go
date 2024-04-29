package registry

import (
	"encoding/json"
	"os"

	"github.com/pernydev/fossdmx/pkg/global"
	"github.com/pernydev/fossdmx/pkg/models"
)

var FixtureTypes = make(map[string]*models.FixtureType)
var FixtureModes = make(map[string]*models.FixtureMode)

func RegisterFixtureType(ft models.FixtureType) {
	FixtureTypes[ft.ID] = &ft
	for _, mode := range ft.Modes {
		RegisterFixtureMode(mode)
	}
}

func RegisterFixtureMode(mode models.FixtureMode) {
	FixtureModes[mode.ID] = &mode
}

func RegisterFixturesPath(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	fixtures := []models.FixtureType{}
	err = decoder.Decode(&fixtures)
	if err != nil {
		return err
	}

	for _, fixture := range fixtures {
		RegisterFixtureType(fixture)
	}

	global.Log.Info("Registered fixtures", "amount", len(fixtures))
	return nil
}
