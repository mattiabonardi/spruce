package models

type Entity struct {
	Class      string
	Attributes map[string]Attribute
}

type Attribute struct {
	Type  PrimitiveType
	Value any
}

type PrimitiveType string

const (
	Integer      PrimitiveType = "Integer"
	String       PrimitiveType = "String"
	Number       PrimitiveType = "Number"
	Boolean      PrimitiveType = "Boolean"
	ObjectId     PrimitiveType = "ObjectId"
	Obect        PrimitiveType = "Object"
	IntegerArray PrimitiveType = "IntegerArray"
	StringArray  PrimitiveType = "StringArray"
	NumberArray  PrimitiveType = "NumberArray"
	ObjectArray  PrimitiveType = "ObjectArray"
)

type EntityContext struct {
	Filters []EntityFilter
}
