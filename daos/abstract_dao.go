package daos

import "github.com/mattiabonardi/spruce/models"

type AbstractDAO interface {
	GetById(executionContext models.ExecutionContext, entityContext models.EntityContext, _id string) (models.Entity, error)
	GetAll(executionContext models.ExecutionContext, entityContext models.EntityContext) ([]models.Entity, error)
	Create(executionContext models.ExecutionContext, entityContext models.EntityContext, entity models.Entity) (models.Entity, error)
	DeleteById(executionContext models.ExecutionContext, entityContext models.EntityContext, _id string) error
}
