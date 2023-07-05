package models

type EntityIterator struct {
	Index    int
	Entities []*Entity
}

func (i *EntityIterator) HasNext() bool {
	return i.Index < len(i.Entities)
}

func (i *EntityIterator) Next() *Entity {
	if i.HasNext() {
		entity := i.Entities[i.Index]
		i.Index++
		return entity
	}
	return nil
}
