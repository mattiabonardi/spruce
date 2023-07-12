package entity

import (
	"bufio"
	"os"

	"github.com/mattiabonardi/spruce/utils"
	"gopkg.in/yaml.v2"
)

type EntityDefinitionManager struct{}

func (h *EntityDefinitionManager) GetDefinition(entityClass string) (EntityDefinition, error) {
	entityDefinition := EntityDefinition{}
	// open file
	filePath := entityClass + ".yaml"
	file, err := os.Open(utils.RootDir() + "/resources/entities/definitions/" + filePath)
	if err != nil {
		return entityDefinition, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var yamlContent string
	for scanner.Scan() {
		yamlContent += scanner.Text() + "\n"
	}
	err = yaml.Unmarshal([]byte(yamlContent), &entityDefinition)
	if err != nil {
		return entityDefinition, err
	}
	return entityDefinition, nil
}
