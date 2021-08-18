package cast

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFromString(t *testing.T) {
	// string

	val, err := FromString("string", "string")
	assert.NoError(t, err)
	assert.Equal(t, "string", val)

	// int

	val, err = FromString("1", "int")
	assert.NoError(t, err)
	assert.Equal(t, 1, val)

	val, err = FromString("str", "int")
	assert.Error(t, err)

	val, err = FromString("1", "int8")
	assert.NoError(t, err)
	assert.Equal(t, int8(1), val)

	val, err = FromString("str", "int8")
	assert.Error(t, err)

	val, err = FromString("1", "int16")
	assert.NoError(t, err)
	assert.Equal(t, int16(1), val)

	val, err = FromString("str", "int16")
	assert.Error(t, err)

	val, err = FromString("1", "int32")
	assert.NoError(t, err)
	assert.Equal(t, int32(1), val)

	val, err = FromString("str", "int32")
	assert.Error(t, err)

	val, err = FromString("1", "int64")
	assert.NoError(t, err)
	assert.Equal(t, int64(1), val)

	val, err = FromString("str", "int64")
	assert.Error(t, err)

	// uint

	val, err = FromString("1", "uint")
	assert.NoError(t, err)
	assert.Equal(t, uint(1), val)

	val, err = FromString("-1", "uint")
	assert.Error(t, err)

	val, err = FromString("1", "uint8")
	assert.NoError(t, err)
	assert.Equal(t, uint8(1), val)

	val, err = FromString("-1", "uint8")
	assert.Error(t, err)

	val, err = FromString("1", "uint16")
	assert.NoError(t, err)
	assert.Equal(t, uint16(1), val)

	val, err = FromString("-1", "uint16")
	assert.Error(t, err)

	val, err = FromString("1", "uint32")
	assert.NoError(t, err)
	assert.Equal(t, uint32(1), val)

	val, err = FromString("-1", "uint32")
	assert.Error(t, err)

	val, err = FromString("1", "uint64")
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), val)

	val, err = FromString("-1", "uint64")
	assert.Error(t, err)

	// float

	val, err = FromString("3.14", "float32")
	assert.NoError(t, err)
	assert.Equal(t, float32(3.14), val)

	val, err = FromString("str", "float32")
	assert.Error(t, err)

	val, err = FromString("3.14", "float64")
	assert.NoError(t, err)
	assert.Equal(t, 3.14, val)

	val, err = FromString("str", "float64")
	assert.Error(t, err)

	// bool

	val, err = FromString("true", "bool")
	assert.NoError(t, err)
	assert.Equal(t, true, val)

	val, err = FromString("1", "bool")
	assert.NoError(t, err)
	assert.Equal(t, true, val)

	val, err = FromString("false", "bool")
	assert.NoError(t, err)
	assert.Equal(t, false, val)

	val, err = FromString("0", "bool")
	assert.NoError(t, err)
	assert.Equal(t, false, val)

	val, err = FromString("invalid", "bool")
	assert.Error(t, err)

	// else

	val, err = FromString("0", "invalid")
	assert.Error(t, err, "cast: type %v is not supported", "invalid")
}
