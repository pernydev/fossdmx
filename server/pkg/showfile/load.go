package showfile

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pernydev/fossdmx/pkg/global"
	"github.com/pernydev/fossdmx/pkg/models"
	"github.com/pernydev/fossdmx/pkg/registry"
)

func LoadFromFile(sfname string) error {
	file, err := os.ReadFile(global.DataFolder + sfname)
	if err != nil {
		return err
	}

	showfile, err := DecodeShowfile(file)
	if err != nil {
		return err
	}

	err = VerifyShowfile(showfile)
	if err != nil {
		return err
	}

	global.Showfile = showfile
	global.ShowfileLoaded = true

	global.Log.Info("Loaded showfile", "name", showfile.Name)

	return nil
}

func DecodeShowfile(file []byte) (models.Showfile, error) {
	showfile := models.Showfile{}
	err := json.Unmarshal(file, &showfile)
	if err != nil {
		return models.Showfile{}, err
	}
	return showfile, nil
}

func VerifyShowfile(showfile models.Showfile) error {
	for _, fixture := range showfile.Fixtures {
		if _, ok := registry.FixtureTypes[fixture.TypeID]; !ok {
			return fmt.Errorf("fixture type %s not found", fixture.TypeID)
		}
		if _, ok := registry.FixtureModes[fixture.ModeID]; !ok {
			return fmt.Errorf("fixture mode %s not found", fixture.ModeID)
		}
	}
	return nil
}

/*
	{
	  "name": "My Showfile",
	  "authors": ["Perny"],
	  "fixtures": {
	    "par_1": {
	      "id": 1,
	      "channel": {
	        "channel": 1,
	        "universe": 1
	      },
	      "type_id": "generic.par",
	      "mode_id": "generic.par.5ch",
	      "visualizor_position": {
	        "x": 0,
	        "y": 0,
	        "z": 0
	      },
	      "overwrite": {
	        "dimmer": 100,
	        "strobe": 0,
	        "color": {
	          "red": 255,
	          "green": 255,
	          "blue": 255
	        }
	      },
	      "layers": []
	    }
	  },
	  "sequences": {},
	  "presets": {}
	}
*/

func CreateDataDirectory() error {
	if _, err := os.Stat(global.DataFolder); os.IsNotExist(err) {
		err = os.Mkdir(global.DataFolder, 0755)
		if err != nil {
			global.Log.Error("Failed to create data folder", "error", err)
			return err
		}
		global.Log.Info("Data folder created", "path", global.DataFolder)
	}
	return nil
}

func InitializeDemoShowfile() error {
	if _, err := os.Stat(global.DataFolder + "demo.json"); err == nil {
		return nil
	}

	showfile := models.Showfile{
		Name:     "Demo Showfile",
		Filename: "demo.json",
		Authors:  []string{"Perny"},
		Fixtures: map[string]*models.Fixture{
			"par_1": {
				ID: "par_1",
				Channel: models.Channel{
					Channel:  1,
					Universe: 1,
				},
				TypeID:             "generic.par",
				ModeID:             "generic.par.5ch",
				VisualizorPosition: models.XYZ{X: 0, Y: 0, Z: 0},
				Layers:             map[string]models.FixtureLayer{},
			},
		},
		Sequences: map[string]*models.Sequence{},
		Scenes:    map[string]*models.Scene{},
	}

	showfilecontent, err := json.Marshal(showfile)
	if err != nil {
		global.Log.Error("Failed to marshal demo showfile", "error", err)
		return err
	}

	err = os.WriteFile(global.DataFolder+"demo.json", showfilecontent, 0644)
	if err != nil {
		global.Log.Error("Failed to save demo showfile", "error", err)
		return err
	}

	global.Log.Info("Saved demo showfile", "name", showfile.Name)
	return nil
}
