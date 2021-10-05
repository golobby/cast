// Package cast is a lightweight casting (type conversion) library.
package cast

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const (
	Uint8  = "uint8"
	Uint16 = "uint16"
	Uint32 = "uint32"
	Uint64 = "uint64"
	Uint   = "uint"

	Int8  = "int8"
	Int16 = "int16"
	Int32 = "int32"
	Int64 = "int64"
	Int   = "int"

	Float32 = "float32"
	Float64 = "float64"

	Bool = "bool"

	String = "string"
)

// FromString casts a string value to the given type name.
func FromString(value string, targetType string) (interface{}, error) {
	message := "cast: cannot cast `%v` to type `%v`"

	switch targetType {
	case Int:
		v, err := strconv.ParseInt(value, 0, 32)
		if err != nil {
			return nil, fmt.Errorf(message, value, targetType)
		}
		return int(v), nil
	case Int8:
		v, err := strconv.ParseInt(value, 0, 8)
		if err != nil {
			return nil, err
		}
		return int8(v), nil
	case Int16:
		v, err := strconv.ParseInt(value, 0, 16)
		if err != nil {
			return nil, err
		}
		return int16(v), nil
	case Int32:
		v, err := strconv.ParseInt(value, 0, 32)
		if err != nil {
			return nil, err
		}
		return int32(v), nil
	case Int64:
		v, err := strconv.ParseInt(value, 0, 64)
		if err != nil {
			return nil, err
		}
		return v, nil

	case Uint:
		v, err := strconv.ParseUint(value, 0, 32)
		if err != nil {
			return nil, err
		}
		return uint(v), nil
	case Uint8:
		v, err := strconv.ParseUint(value, 0, 8)
		if err != nil {
			return nil, err
		}
		return uint8(v), nil
	case Uint16:
		v, err := strconv.ParseUint(value, 0, 16)
		if err != nil {
			return nil, err
		}
		return uint16(v), nil
	case Uint32:
		v, err := strconv.ParseUint(value, 0, 32)
		if err != nil {
			return nil, err
		}
		return uint32(v), nil
	case Uint64:
		v, err := strconv.ParseUint(value, 0, 64)
		if err != nil {
			return nil, err
		}
		return v, nil

	case Bool:
		v, err := strconv.ParseBool(value)
		if err != nil {
			return nil, err
		}
		return v, nil

	case Float32:
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}
		return float32(v), nil
	case Float64:
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}
		return v, nil

	case String:
		return value, nil
	}

	return nil, fmt.Errorf("cast: type %v is not supported", targetType)
}

// FromType casts a string value to the given reflected type.
func FromType(value string, targetType reflect.Type) (interface{}, error) {
	var typeName = targetType.String()

	if strings.HasPrefix(typeName, "[]") {
		itemType := typeName[2:]
		array := reflect.New(targetType).Elem()

		for _, v := range strings.Split(value, ",") {
			if item, err := FromString(strings.Trim(v, " \n\r"), itemType); err != nil {
				return array.Interface(), err
			} else {
				array = reflect.Append(array, reflect.ValueOf(item))
			}
		}

		return array.Interface(), nil
	}

	return FromString(value, typeName)
}
