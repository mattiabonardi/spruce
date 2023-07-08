package daos

import (
	"bufio"
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

func (h YamlDAO) GetAll(executionContext models.ExecutionContext, entityContext models.EntityContext) ([]models.Entity, error) {
	var entities = []models.Entity{}
	// get yaml data
	entityData, err := h.getYamlData()
	if err != nil {
		return entities, err
	}
	for _, e := range entityData.Data {
		entities = append(entities, h.buildEntityFromYamlRecord(e))
	}
	return entities, nil
}

func (h YamlDAO) getYamlData() (YamlEntityData, error) {
	var yamlEntityData = YamlEntityData{}
	// open file
	file, err := os.Open(utils.RootDir() + "/" + h.EntityDefinition.DataSource.YamlFileConfig.FilePath)
	if err != nil {
		return yamlEntityData, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var yamlContent string
	for scanner.Scan() {
		yamlContent += scanner.Text() + "\n"
	}
	err = yaml.Unmarshal([]byte(yamlContent), &yamlEntityData)
	if err != nil {
		return yamlEntityData, err
	}
	return yamlEntityData, nil
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
