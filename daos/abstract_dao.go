package daos

import "github.com/mattiabonardi/spruce/models"

type AbstractDAO interface {
	//GetById(executionContext models.ExecutionContext, entityContext models.EntityContext, id string) models.Entity
	GetAll(executionContext models.ExecutionContext, entityContext models.EntityContext) []models.Entity
}
