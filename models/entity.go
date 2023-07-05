package models

type Entity struct {
	Name        string
	Description string
	Attributes  Attributes
}

type Attributes struct {
	Instances   map[string]Attribute
	Id          string
	Description string
}

type Attribute struct {
	Type PrimitiveType
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
