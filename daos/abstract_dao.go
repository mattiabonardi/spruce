package daos

import "github.com/mattiabonardi/spruce/models"

type AbstractDAO interface {
	GetById(executionContext models.ExecutionContext, entityContext models.EntityContext, _id string) (models.Entity, error)
	GetAll(executionContext models.ExecutionContext, entityContext models.EntityContext) ([]models.Entity, error)
}
