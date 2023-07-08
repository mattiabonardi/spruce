package models

type EntityDefinition struct {
	Class       string                         `yaml:"class"`
	Description string                         `yaml:"description"`
	DataSource  DataSource                     `yaml:"dataSource"`
	Attributes  map[string]AttributeDefinition `yaml:"attributes"`
}

type AttributeDefinition struct {
	Type Type `yaml:"type"`
}

type DataSourceType string

const (
	YamlFile DataSourceType = "YamlFile"
)

type DataSource struct {
	Type           DataSourceType `yaml:"type"`
	YamlFileConfig YamlDAOConfig  `yaml:"yamlFileConfig"`
}

type YamlDAOConfig struct {
	FilePath string `yaml:"filePath"`
}
