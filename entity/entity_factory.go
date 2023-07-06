package entity

import (
	"log"

	"github.com/mattiabonardi/spruce/daos"
	"github.com/mattiabonardi/spruce/models"
)

type EntityFactory struct{}

func (h EntityFactory) CreateDAO(entityDefinition models.EntityDefinition) daos.AbstractDAO {
	switch entityDefinition.DataSource.Type {
	case models.YamlFile:
		return daos.YamlDAO{
			EntityDefinition: entityDefinition,
		}
	default:
		log.Fatal("DAO not supported")
	}
	return nil
}
