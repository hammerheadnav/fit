package fit

import "reflect"

// DeveloperField holds the value of a developer field and associated metadata to allow for dynamic parsing.
type DeveloperField struct {
	DeveloperDataIndex    uint8
	FieldDefinitionNumber uint8
	BaseTypeId            uint8
	FieldName             string
	Units                 string
	Value                 reflect.Value
}
