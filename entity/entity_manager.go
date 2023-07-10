package entity

import "github.com/mattiabonardi/spruce/models"

type EntityManager struct {
	ExecutionContext models.ExecutionContext
}

func (h EntityManager) GetById(entityClass string, entityContext models.EntityContext, _id string) (models.Entity, error) {
	// get entity definition
	entityDefinitionManager := EntityDefinitionManager{}
	entityDefinition, err := entityDefinitionManager.GetDefinition(entityClass)
	if err != nil {
		return models.Entity{}, err
	}
	// call GetAll
	return EntityFactory{}.CreateDAO(entityDefinition).GetById(h.ExecutionContext, entityContext, _id)
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

func (h EntityManager) Create(entityClass string, entityContext models.EntityContext, entity models.Entity) (models.Entity, error) {
	// get entity definition
	entityDefinitionManager := EntityDefinitionManager{}
	entityDefinition, err := entityDefinitionManager.GetDefinition(entityClass)
	if err != nil {
		return models.Entity{}, err
	}
	// call Create
	return EntityFactory{}.CreateDAO(entityDefinition).Create(h.ExecutionContext, entityContext, entity)
}

func (h EntityManager) DeleteById(entityClass string, entityContext models.EntityContext, _id string) error {
	// get entity definition
	entityDefinitionManager := EntityDefinitionManager{}
	entityDefinition, err := entityDefinitionManager.GetDefinition(entityClass)
	if err != nil {
		return err
	}
	// call DeleteById
	return EntityFactory{}.CreateDAO(entityDefinition).DeleteById(h.ExecutionContext, entityContext, _id)
}
