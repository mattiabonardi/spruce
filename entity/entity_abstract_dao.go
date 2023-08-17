package entity

import "github.com/mattiabonardi/endor-sdk-go/models"

type AbstractDAO interface {
	GetById(executionContext models.ExecutionContext, entityContext EntityContext, _id string) (Entity, error)
	GetAll(executionContext models.ExecutionContext, entityContext EntityContext) ([]Entity, error)
	Create(executionContext models.ExecutionContext, entityContext EntityContext, entity Entity) (Entity, error)
	DeleteById(executionContext models.ExecutionContext, entityContext EntityContext, _id string) error
	Update(executionContext models.ExecutionContext, entityContext EntityContext, entity Entity) error
}
