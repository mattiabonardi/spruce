package entity

import (
	"testing"

	"github.com/mattiabonardi/spruce/models"
)

func TestGetDefinition(t *testing.T) {
	entityDefinitionManager := EntityDefinitionManager{}
	entityDefinition, err := entityDefinitionManager.GetDefinition("gender")

	if err != nil {
		t.Fatal(err)
	}

	if entityDefinition.DataSource.Type != models.YamlFile {
		t.Fatal("Aspected: " + models.YamlFile + " got: " + entityDefinition.DataSource.Type)
	}
}
