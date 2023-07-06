package entity

import (
	"testing"

	"github.com/mattiabonardi/spruce/models"
	"gopkg.in/yaml.v2"
)

func GetDefinitionTest(t *testing.T) {
	entityDefinition := EntityDefinitionManager{}.GetDefinition("gender")

	t.Log(yaml.Marshal(entityDefinition))

	if entityDefinition.DataSource.Type != "aaaa" {
		t.Fatal("Aspected: " + models.YamlDAO + " got: " + entityDefinition.DataSource.Type)
	}
}
