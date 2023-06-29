package models

type Entity struct {
	name        string
	description string
	attributes  Attributes
}

type Attributes struct {
	instances   map[string]Attribute
	id          string
	description string
}

type Attribute struct {
	_type PrimitiveType
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
