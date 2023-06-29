package models

type EntityDefinition struct {
	name        string               `yaml:"name"`
	description string               `yaml:"description"`
	attributes  AttributesDefinition `yaml:"attributes"`
}

type AttributesDefinition struct {
	instances   map[string]AttributeDefinition `yaml:"instances"`
	id          string                         `yaml:"id"`
	description string                         `yaml:"description"`
}

type AttributeDefinition struct {
	_type PrimitiveType `yaml:"type"`
}
