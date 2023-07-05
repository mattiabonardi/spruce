package models

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
