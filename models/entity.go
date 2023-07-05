package models

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
	IntegerArray Type = "IntegerArray"
	StringArray  Type = "StringArray"
	NumberArray  Type = "NumberArray"
	ObjectArray  Type = "ObjectArray"
)

type EntityContext struct {
	Filters []EntityFilter
}
