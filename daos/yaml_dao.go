package daos

import (
	"io/ioutil"
	"log"

	"github.com/mattiabonardi/spruse/models"
	"github.com/mattiabonardi/spruse/utils"
	"gopkg.in/yaml.v2"
)

var yamlEntityDataPath string = utils.RootDir() + "resources/data/"

type YamlEntityData struct {
	Data []map[string]interface{} `yaml:"data"`
}

type YamlDAO struct {
	EntityDefinition models.EntityDefinition
}

func (h YamlDAO) GetAll(executionContext models.ExecutionContext, entityContext models.EntityContext) models.EntityIterator {
	// get file name
	fileName := h.EntityDefinition.Name + ".yaml"
	// read yaml file
	yamlFile, err := ioutil.ReadFile(yamlEntityDataPath + fileName)
	if err != nil {
		log.Fatalf("Unable to read "+fileName+" %v", err)
	}
	// unmarshal yaml file to YamlEntityData
	entityData := YamlEntityData{}
	err = yaml.Unmarshal(yamlFile, entityData)
	if err != nil {
		log.Fatalf(fileName+" format not valid: %v", err)
	}
	var entities = []models.Entity{}
	for i, e := range entityData.Data {
		entities[i] = buildEntityFromYamlRecord(h.EntityDefinition, e)
	}
	var EntityIterator = []models.Entity{}
	EntityIterator.
}

func buildEntityFromYamlRecord (entityDefinition models.EntityDefinition, yamlRecord map[string]interface{}) models.Entity {

}

/*func (h YamlDAO) GetById(executionContext models.ExecutionContext, entityContext models.EntityContext, id string) models.Entity {

}*/
