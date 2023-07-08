package daos

import (
	"bufio"
	"log"
	"os"

	"github.com/mattiabonardi/spruce/models"
	"github.com/mattiabonardi/spruce/utils"
	"gopkg.in/yaml.v2"
)

type YamlEntityData struct {
	Data []map[string]interface{} `yaml:"data"`
}

type YamlDAO struct {
	EntityDefinition models.EntityDefinition
}

func (h YamlDAO) GetAll(executionContext models.ExecutionContext, entityContext models.EntityContext) []models.Entity {
	// get yaml data
	entityData := h.getYamlData()
	var entities = []models.Entity{}
	for _, e := range entityData.Data {
		entities = append(entities, h.buildEntityFromYamlRecord(e))
	}
	return entities
}

func (h YamlDAO) getYamlData() YamlEntityData {
	// open file
	file, err := os.Open(utils.RootDir() + "/" + h.EntityDefinition.DataSource.YamlFileConfig.FilePath)
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var yamlContent string
	for scanner.Scan() {
		yamlContent += scanner.Text() + "\n"
	}
	var YamlEntityData = YamlEntityData{}
	err = yaml.Unmarshal([]byte(yamlContent), &YamlEntityData)
	if err != nil {
		log.Fatalf("Unable to elaborate file: %v", err)
	}
	return YamlEntityData
}

func (h YamlDAO) buildEntityFromYamlRecord(yamlRecord map[string]interface{}) models.Entity {
	var entity = models.Entity{}
	entity.Class = h.EntityDefinition.Class
	var attributes = make(map[string]models.Attribute)
	// iterate definitions
	for k, v := range h.EntityDefinition.Attributes {
		// create attribute
		attribute := models.Attribute{}
		attribute.Type = v.Type
		attribute.Value = yamlRecord[k]
		attributes[k] = attribute
	}
	entity.Attributes = attributes
	return entity
}
