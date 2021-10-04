// Package cast is a lightweight casting (type conversion) library.
package cast

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	Uint8  = "uint8"
	Uint16 = "uint16"
	Uint32 = "uint32"
	Uint64 = "uint64"
	Uint   = "uint"

	Uint8Array  = "[]uint8"
	Uint16Array = "[]uint16"
	Uint32Array = "[]uint32"
	Uint64Array = "[]uint64"
	UintArray   = "[]uint"

	Int8  = "int8"
	Int16 = "int16"
	Int32 = "int32"
	Int64 = "int64"
	Int   = "int"

	Int8Array  = "[]int8"
	Int16Array = "[]int16"
	Int32Array = "[]int32"
	Int64Array = "[]int64"
	IntArray   = "[]int"

	Float32 = "float32"
	Float64 = "float64"

	Float32Array = "[]float32"
	Float64Array = "[]float64"

	Bool = "bool"

	BoolArray = "[]bool"

	String = "string"

	StringArray = "[]string"
)

// FromString casts a string value to the given type.
func FromString(value string, targetType string) (interface{}, error) {
	var message = "cast: cannot cast `%v` to type `%v`"

	if strings.HasPrefix(targetType, "[]") {
		kind := targetType[2:]

		switch kind {
		case Int:
			var a []int
			if err := Map(value, kind, func(v interface{}) { a = append(a, v.(int)) }); err != nil {
				return nil, err
			}
			return a, nil
		case Int8:
			var a []int8
			if err := Map(value, kind, func(v interface{}) { a = append(a, v.(int8)) }); err != nil {
				return nil, err
			}
			return a, nil
		case Int16:
			var a []int16
			if err := Map(value, kind, func(v interface{}) { a = append(a, v.(int16)) }); err != nil {
				return nil, err
			}
			return a, nil
		case Int32:
			var a []int32
			if err := Map(value, kind, func(v interface{}) { a = append(a, v.(int32)) }); err != nil {
				return nil, err
			}
			return a, nil
		case Int64:
			var a []int64
			if err := Map(value, kind, func(v interface{}) { a = append(a, v.(int64)) }); err != nil {
				return nil, err
			}
			return a, nil

		case Uint:
			var a []uint
			if err := Map(value, kind, func(v interface{}) { a = append(a, v.(uint)) }); err != nil {
				return nil, err
			}
			return a, nil
		case Uint8:
			var a []uint8
			if err := Map(value, kind, func(v interface{}) { a = append(a, v.(uint8)) }); err != nil {
				return nil, err
			}
			return a, nil
		case Uint16:
			var a []uint16
			if err := Map(value, kind, func(v interface{}) { a = append(a, v.(uint16)) }); err != nil {
				return nil, err
			}
			return a, nil
		case Uint32:
			var a []uint32
			if err := Map(value, kind, func(v interface{}) { a = append(a, v.(uint32)) }); err != nil {
				return nil, err
			}
			return a, nil
		case Uint64:
			var a []uint64
			if err := Map(value, kind, func(v interface{}) { a = append(a, v.(uint64)) }); err != nil {
				return nil, err
			}
			return a, nil

		case Float32:
			var a []float32
			if err := Map(value, kind, func(v interface{}) { a = append(a, v.(float32)) }); err != nil {
				return nil, err
			}
			return a, nil
		case Float64:
			var a []float64
			if err := Map(value, kind, func(v interface{}) { a = append(a, v.(float64)) }); err != nil {
				return nil, err
			}
			return a, nil

		case Bool:
			var a []bool
			if err := Map(value, kind, func(v interface{}) { a = append(a, v.(bool)) }); err != nil {
				return nil, err
			}
			return a, nil

		case String:
			var a []string
			if err := Map(value, kind, func(v interface{}) { a = append(a, v.(string)) }); err != nil {
				return nil, err
			}
			return a, nil

		default:
			return nil, fmt.Errorf(message, value, targetType)
		}
	}

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

	case Bool:
		v, err := strconv.ParseBool(value)
		if err != nil {
			return nil, err
		}
		return v, nil

	case String:
		return value, nil
	}

	return nil, fmt.Errorf("cast: type %v is not supported", targetType)
}

func Map(csv, kind string, callback func(v interface{})) error {
	for _, value := range strings.Split(csv, ",") {
		if typedValue, err := FromString(value, kind); err != nil {
			return err
		} else {
			callback(typedValue)
		}
	}

	return nil
}
