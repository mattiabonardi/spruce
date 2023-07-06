package entity

import (
	"testing"

	"github.com/mattiabonardi/spruce/models"
)

func TestGetDefinition(t *testing.T) {
	entityDefinitionManager := EntityDefinitionManager{}
	entityDefinition := entityDefinitionManager.GetDefinition("gender")

	if entityDefinition.DataSource.Type != models.YamlFile {
		t.Fatal("Aspected: " + models.YamlFile + " got: " + entityDefinition.DataSource.Type)
	}
}
