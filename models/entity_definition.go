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
	YamlDAO DataSourceType = "YamlDAO"
)

type DataSource struct {
	Type          DataSourceType `yaml:"type"`
	YamlDAOConfig YamlDAOConfig  `yaml:"yamlDAOConfig"`
}

type YamlDAOConfig struct {
	FilePath string `yaml:"type"`
}
