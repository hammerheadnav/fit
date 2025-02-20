package fit

import (
	"math"
	"reflect"
)

// DeveloperField holds the value of a developer field and associated metadata to allow for dynamic parsing.
type DeveloperField struct {
	DeveloperDataIndex    uint8
	FieldDefinitionNumber uint8
	BaseTypeId            FitBaseType
	FieldName             string
	Units                 string
	value                 reflect.Value
}

func first(v reflect.Value) reflect.Value {
	switch v.Kind() {
	case reflect.Slice,
		reflect.Array:
		if v.Len() > 0 {
			return v.Index(0)
		}
	}
	return v
}

func getUint64(v reflect.Value) uint64 {
	switch v.Kind() {
	case reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		return v.Uint()
	default:
		return 0xFFFFFFFFFFFFFFFF
	}
}

func getFloat64(v reflect.Value) float64 {
	switch v.Kind() {
	case reflect.Float32,
		reflect.Float64:
		return v.Float()
	default:
		return math.NaN()
	}
}

func (f *DeveloperField) Uint64() uint64 {
	return getUint64(first(f.value))
}

func (f *DeveloperField) Float64() float64 {
	return getFloat64(first(f.value))
}

func (f *DeveloperField) Uint64Slice() []uint64 {
	var values []uint64

	switch f.value.Kind() {
	case reflect.Slice:
		for i := 0; i < f.value.Len(); i++ {
			v := getUint64(f.value.Index(i))
			values = append(values, v)
		}
	default:
		v := getUint64(f.value)
		if v != 0xFFFFFFFFFFFFFFFF {
			values = append(values, v)
		}
	}

	return values
}
