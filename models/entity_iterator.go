package models

type EntityIterator struct {
	index    int
	entities []*Entity
}

func (i *EntityIterator) hasNext() bool {
	if i.index < len(i.entities) {
		return true
	}
	return false
}

func (i *EntityIterator) next() *Entity {
	if i.hasNext() {
		entity := i.entities[i.index]
		i.index++
		return entity
	}
	return nil
}
