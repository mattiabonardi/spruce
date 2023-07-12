package entity

import (
	"log"
)

type EntityFactory struct{}

func (h EntityFactory) CreateDAO(entityDefinition EntityDefinition) AbstractDAO {
	switch entityDefinition.DataSource.Type {
	case YamlFile:
		return YamlDAO{
			EntityDefinition: entityDefinition,
		}
	default:
		log.Fatal("DAO not supported")
	}
	return nil
}
