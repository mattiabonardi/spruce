package entity

import (
	"bufio"
	"log"
	"os"

	"github.com/mattiabonardi/spruce/models"
	"github.com/mattiabonardi/spruce/utils"
	"gopkg.in/yaml.v2"
)

type EntityDefinitionManager struct{}

var definitionBasePath = utils.RootDir() + "/resources/entities/definitions/"

func (h EntityDefinitionManager) GetDefinition(entityClass string) models.EntityDefinition {
	// open file
	filePath := entityClass + ".yaml"
	file, err := os.Open(definitionBasePath + "/" + filePath)
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var yamlContent string
	for scanner.Scan() {
		yamlContent += scanner.Text() + "\n"
	}
	entityDefinition := models.EntityDefinition{}
	err = yaml.Unmarshal([]byte(yamlContent), &entityDefinition)
	if err != nil {
		log.Fatalf("Unable to elaborate file: %v", err)
	}
	return entityDefinition
}
