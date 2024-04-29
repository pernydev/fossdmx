package fixtures

import (
	"github.com/pernydev/fossdmx/pkg/models"
	"github.com/pernydev/fossdmx/pkg/registry"
)

func UpdateFixture(fixture *models.Fixture) {
	registry.FixtureModes[fixture.ModeID].Update(fixture)
}
