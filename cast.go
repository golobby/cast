// Package cast is a lightweight yet powerful casting library for Go projects.
package cast

import (
	"fmt"
	"strconv"
)

// FromString casts from a string variable to the given type
func FromString(value string, t string) (interface{}, error) {
	switch t {
	case "int":
		v, err := strconv.ParseInt(value, 0, 64)
		if err != nil {
			return nil, err
		}
		return int(v), nil
	case "int8":
		v, err := strconv.ParseInt(value, 0, 8)
		if err != nil {
			return nil, err
		}
		return int8(v), nil
	case "int16":
		v, err := strconv.ParseInt(value, 0, 16)
		if err != nil {
			return nil, err
		}
		return int16(v), nil
	case "int32":
		v, err := strconv.ParseInt(value, 0, 32)
		if err != nil {
			return nil, err
		}
		return int32(v), nil
	case "int64":
		v, err := strconv.ParseInt(value, 0, 64)
		if err != nil {
			return nil, err
		}
		return v, nil

	case "uint":
		v, err := strconv.ParseUint(value, 0, 64)
		if err != nil {
			return nil, err
		}
		return uint(v), nil
	case "uint8":
		v, err := strconv.ParseUint(value, 0, 8)
		if err != nil {
			return nil, err
		}
		return uint8(v), nil
	case "uint16":
		v, err := strconv.ParseUint(value, 0, 16)
		if err != nil {
			return nil, err
		}
		return uint16(v), nil
	case "uint32":
		v, err := strconv.ParseUint(value, 0, 32)
		if err != nil {
			return nil, err
		}
		return uint32(v), nil
	case "uint64":
		v, err := strconv.ParseUint(value, 0, 64)
		if err != nil {
			return nil, err
		}
		return v, nil

	case "bool":
		v, err := strconv.ParseBool(value)
		if err != nil {
			return nil, err
		}
		return v, nil

	case "float32":
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}
		return float32(v), nil
	case "float64":
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}
		return v, nil

	case "string":
		return value, nil
	}

	return nil, fmt.Errorf("cast: type %v is not supported", t)
}
