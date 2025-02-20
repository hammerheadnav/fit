package fit

import "reflect"

// DeveloperField holds the value of a developer field and associated metadata to allow for dynamic parsing.
type DeveloperField struct {
	DeveloperDataIndex    uint8
	FieldDefinitionNumber uint8
	BaseTypeId            FitBaseType
	FieldName             string
	Units                 string
	value                 interface{}
}

func (f *DeveloperField) Value() interface{} {
	v := reflect.ValueOf(f.value)
	switch v.Type().Kind() {
	case reflect.Slice, reflect.Array:
		return v.Index(0).Interface()
	default:
		return f.value
	}
}

func (f *DeveloperField) Values() []interface{} {
	if f.value == nil {
		return nil
	}

	v := reflect.ValueOf(f.value)
	switch v.Type().Kind() {
	case reflect.Slice, reflect.Array:
		values, ok := (f.value).([]interface{})
		if ok {
			return values
		}

		for i := 0; i < v.Len(); i++ {
			values = append(values, v.Index(i))
		}
		return values
	default:
		return []interface{}{f.value}
	}
}
