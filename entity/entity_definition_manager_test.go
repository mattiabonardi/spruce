package entity

import (
	"testing"
)

func TestGetDefinition(t *testing.T) {
	entityDefinitionManager := EntityDefinitionManager{}
	entityDefinition, err := entityDefinitionManager.GetDefinition("gender")

	if err != nil {
		t.Fatal(err)
	}

	if entityDefinition.DataSource.Type != YamlFile {
		t.Fatal("Aspected: " + YamlFile + " got: " + entityDefinition.DataSource.Type)
	}
}
