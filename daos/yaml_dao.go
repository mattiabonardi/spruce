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
	entityData := getYamlData(h.EntityDefinition)
	var entities = []models.Entity{}
	for i, e := range entityData.Data {
		entities[i] = buildEntityFromYamlRecord(h.EntityDefinition, e)
	}
	return entities
}

func getYamlData(entityDefinition models.EntityDefinition) YamlEntityData {
	// open file
	file, err := os.Open(utils.RootDir() + "/" + entityDefinition.DataSource.YamlFileConfig.FilePath)
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

func buildEntityFromYamlRecord(entityDefinition models.EntityDefinition, yamlRecord map[string]interface{}) models.Entity {
	var entity = models.Entity{}
	entity.Class = entityDefinition.Class
	var attributes = make(map[string]models.Attribute)
	// iterate definitions
	for k, v := range entityDefinition.Attributes {
		// create attribute
		attribute := models.Attribute{}
		attribute.Type = v.Type
		attribute.Value = yamlRecord[k]
		attributes[k] = attribute
	}
	entity.Attributes = attributes
	return entity
}
