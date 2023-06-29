package daos

type AbstractDAO interface {
	getById(id string) Entity
}
