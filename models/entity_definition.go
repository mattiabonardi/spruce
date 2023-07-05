package models

type EntityDefinition struct {
	Class       string                         `yaml:"class"`
	Description string                         `yaml:"description"`
	Attributes  map[string]AttributeDefinition `yaml:"attributes"`
}

type AttributeDefinition struct {
	Type PrimitiveType `yaml:"type"`
}
