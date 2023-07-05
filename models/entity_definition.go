package models

type EntityDefinition struct {
	Name        string               `yaml:"name"`
	Description string               `yaml:"description"`
	Attributes  AttributesDefinition `yaml:"attributes"`
}

type AttributesDefinition struct {
	Instances   map[string]AttributeDefinition `yaml:"instances"`
	Id          string                         `yaml:"id"`
	Description string                         `yaml:"description"`
}

type AttributeDefinition struct {
	Type PrimitiveType `yaml:"type"`
}
