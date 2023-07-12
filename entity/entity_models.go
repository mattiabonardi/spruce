package entity

type Entity struct {
	Class      string
	Attributes map[string]Attribute
}

type Attribute struct {
	Type  Type
	Value any
}

type Type string

const (
	Integer      Type = "Integer"
	String       Type = "String"
	Decimal      Type = "Decimal"
	Boolean      Type = "Boolean"
	ObjectId     Type = "ObjectId"
	Object       Type = "Object"
	IntegerArray Type = "Array<Integer>"
	StringArray  Type = "Array<String>"
	DecimalArray Type = "Array<Decimal>"
	ObjectArray  Type = "Array<Object>"
)

type EntityContext struct {
	Filters []EntityFilter
}

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

type EntityFilter struct {
	Operation FilterOperation
	Attribute string
	Value     any
}

type FilterOperation string

const (
	Equals          FilterOperation = "eq"
	NotEquals       FilterOperation = "ne"
	GreaterThan     FilterOperation = "gt"
	LessThan        FilterOperation = "lt"
	greaterOrEquals FilterOperation = "gte"
	LessOrEquals    FilterOperation = "lte"
	Contains        FilterOperation = "lk"
)
