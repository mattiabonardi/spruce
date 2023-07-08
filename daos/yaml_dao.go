package daos

import (
	"bufio"
	"errors"
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

func (h YamlDAO) GetById(executionContext models.ExecutionContext, entityContext models.EntityContext, _id string) (models.Entity, error) {
	var entity = models.Entity{}
	// get yaml data
	entityData, err := h.getYamlData()
	if err != nil {
		return entity, err
	}
	// iter over elements
	for _, e := range entityData.Data {
		// get _id field as string
		id := e["_id"].(string)
		if id == _id {
			// return entity
			return h.buildEntityFromYamlRecord(e), nil
		}
	}
	return entity, errors.New("No entity matched with _id: " + _id)
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
