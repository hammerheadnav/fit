package fit

// DeveloperField holds the value of a developer field and associated metadata to allow for dynamic parsing.
type DeveloperField struct {
	DeveloperDataIndex    uint8
	FieldDefinitionNumber uint8
	BaseTypeId            FitBaseType
	FieldName             string
	Units                 string
	Value                 interface{}
}
