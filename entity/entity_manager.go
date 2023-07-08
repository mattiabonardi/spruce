package entity

import "github.com/mattiabonardi/spruce/models"

type EntityManager struct {
	ExecutionContext models.ExecutionContext
}

func (h EntityManager) GetAll(entityClass string, entityContext models.EntityContext) ([]models.Entity, error) {
	// get entity definition
	entityDefinitionManager := EntityDefinitionManager{}
	entityDefinition, err := entityDefinitionManager.GetDefinition(entityClass)
	if err != nil {
		return nil, err
	}
	// call GetAll
	return EntityFactory{}.CreateDAO(entityDefinition).GetAll(h.ExecutionContext, entityContext)
}
