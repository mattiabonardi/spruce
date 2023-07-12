package entity

import "github.com/mattiabonardi/spruce/execution"

type AbstractDAO interface {
	GetById(executionContext execution.ExecutionContext, entityContext EntityContext, _id string) (Entity, error)
	GetAll(executionContext execution.ExecutionContext, entityContext EntityContext) ([]Entity, error)
	Create(executionContext execution.ExecutionContext, entityContext EntityContext, entity Entity) (Entity, error)
	DeleteById(executionContext execution.ExecutionContext, entityContext EntityContext, _id string) error
	Update(executionContext execution.ExecutionContext, entityContext EntityContext, entity Entity) error
}
