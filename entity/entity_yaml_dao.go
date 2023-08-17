package entity

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/mattiabonardi/endor-sdk-go/models"
	"gopkg.in/yaml.v3"
)

type YamlEntityData struct {
	Data []map[string]interface{} `yaml:"data"`
}

type YamlDAO struct {
	EntityDefinition EntityDefinition
}

func (h YamlDAO) GetById(executionContext models.ExecutionContext, entityContext EntityContext, _id string) (Entity, error) {
	var entity = Entity{}
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
	return entity, fmt.Errorf("no entity matched with _id: %s", _id)
}

func (h YamlDAO) GetAll(executionContext models.ExecutionContext, entityContext EntityContext) ([]Entity, error) {
	var entities = []Entity{}
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

func (h YamlDAO) Create(executionContext models.ExecutionContext, entityContext EntityContext, entity Entity) (Entity, error) {
	// get yaml data
	entityData, err := h.getYamlData()
	if err != nil {
		return entity, err
	}
	// check if already exist
	for _, e := range entityData.Data {
		if e["_id"] == entity.Attributes["_id"].Value {
			return entity, fmt.Errorf("entity with id: %s already exist", entity.Attributes["_id"])
		}
	}
	// create new item in entityData
	item := make(map[string]interface{})
	for k := range h.EntityDefinition.Attributes {
		item[k] = entity.Attributes[k].Value
	}
	// add to list
	entityData.Data = append(entityData.Data, item)
	// write file
	err = h.saveYamlData(entityData)
	if err != nil {
		return entity, err
	}
	return entity, nil
}

func (h YamlDAO) DeleteById(executionContext models.ExecutionContext, entityContext EntityContext, _id string) error {
	// get yaml data
	entityData, err := h.getYamlData()
	if err != nil {
		return err
	}
	// delete element from yaml data
	var index = -1
	for i, e := range entityData.Data {
		if e["_id"] == _id {
			index = i
		}
	}
	if index != -1 {
		copy(entityData.Data[index:], entityData.Data[index+1:])
		entityData.Data = entityData.Data[:len(entityData.Data)-1]
	}
	// write file
	err = h.saveYamlData(entityData)
	if err != nil {
		return err
	}
	return nil
}

func (h YamlDAO) Update(executionContext models.ExecutionContext, entityContext EntityContext, entity Entity) error {
	// get yaml data
	entityData, err := h.getYamlData()
	if err != nil {
		return err
	}
	// find entity with same id
	var index = -1
	for i, e := range entityData.Data {
		if e["_id"] == entity.Attributes["_id"].Value {
			index = i
		}
	}
	if index == -1 {
		return fmt.Errorf("entity with id %s not found", entity.Attributes["_id"].Value)
	}
	// merge old data with new data
	e := entityData.Data[index]
	for name, value := range entity.Attributes {
		e[name] = value.Value
	}
	// set new entity
	entityData.Data[index] = e
	// write file
	err = h.saveYamlData(entityData)
	if err != nil {
		return err
	}
	return nil
}

func (h YamlDAO) getYamlData() (YamlEntityData, error) {
	var yamlEntityData = YamlEntityData{}
	// open file
	file, err := os.Open(h.getYamlPath())
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

func (h YamlDAO) saveYamlData(yamlEntityData YamlEntityData) error {
	// convert to string yaml
	yamlContent, err := yaml.Marshal(yamlEntityData)
	if err != nil {
		return err
	}
	// open file
	file, err := os.OpenFile(h.getYamlPath(), os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	// write file
	_, err = file.Write(yamlContent)
	if err != nil {
		return err
	}
	return nil
}

func (h YamlDAO) getYamlPath() string {
	return h.rootDir() + "/" + h.EntityDefinition.DataSource.YamlFileConfig.FilePath
}

func (h YamlDAO) buildEntityFromYamlRecord(yamlRecord map[string]interface{}) Entity {
	var entity = Entity{}
	entity.Class = h.EntityDefinition.Class
	var attributes = make(map[string]Attribute)
	// iterate definitions
	for k, v := range h.EntityDefinition.Attributes {
		// create attribute
		attribute := Attribute{}
		attribute.Type = v.Type
		attribute.Value = yamlRecord[k]
		attributes[k] = attribute
	}
	entity.Attributes = attributes
	return entity
}

func (h YamlDAO) rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
