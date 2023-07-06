package entity

import (
	"fmt"
	"testing"

	"github.com/mattiabonardi/spruce/models"
	"gopkg.in/yaml.v2"
)

func TestGetDefinition(t *testing.T) {
	entityDefinitionManager := EntityDefinitionManager{}
	entityDefinition := entityDefinitionManager.GetDefinition("gender")

	yamlData, _ := yaml.Marshal(entityDefinition)
	fmt.Println(string(yamlData))

	if entityDefinition.DataSource.Type != models.YamlFile {
		t.Fatal("Aspected: " + models.YamlFile + " got: " + entityDefinition.DataSource.Type)
	}
}
