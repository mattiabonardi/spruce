package entity

import "github.com/mattiabonardi/spruce/models"

type EntityManager struct {
	ExecutionContext models.ExecutionContext
}

func (h EntityManager) GetAll(entityClass string, entityContext models.EntityContext) []models.Entity {
	// get entity definition
	entityDefinitionManager := EntityDefinitionManager{}
	entityDefinition := entityDefinitionManager.GetDefinition(entityClass)
	// call GetAll
	return EntityFactory{}.CreateDAO(entityDefinition).GetAll(h.ExecutionContext, entityContext)
}
