package cast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromString(t *testing.T) {
	// string

	val, err := FromString("string", String)
	assert.NoError(t, err)
	assert.Equal(t, "string", val)

	// int

	val, err = FromString("1", Int)
	assert.NoError(t, err)
	assert.Equal(t, 1, val)

	_, err = FromString("str", Int)
	assert.Error(t, err)

	val, err = FromString("1", Int8)
	assert.NoError(t, err)
	assert.Equal(t, int8(1), val)

	_, err = FromString("str", Int8)
	assert.Error(t, err)

	val, err = FromString("1", Int16)
	assert.NoError(t, err)
	assert.Equal(t, int16(1), val)

	_, err = FromString("str", Int16)
	assert.Error(t, err)

	val, err = FromString("1", Int32)
	assert.NoError(t, err)
	assert.Equal(t, int32(1), val)

	_, err = FromString("str", Int32)
	assert.Error(t, err)

	val, err = FromString("1", Int64)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), val)

	_, err = FromString("str", Int64)
	assert.Error(t, err)

	// uint

	val, err = FromString("1", Uint)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), val)

	_, err = FromString("-1", Uint)
	assert.Error(t, err)

	val, err = FromString("1", Uint8)
	assert.NoError(t, err)
	assert.Equal(t, uint8(1), val)

	_, err = FromString("-1", Uint8)
	assert.Error(t, err)

	val, err = FromString("1", Uint16)
	assert.NoError(t, err)
	assert.Equal(t, uint16(1), val)

	_, err = FromString("-1", Uint16)
	assert.Error(t, err)

	val, err = FromString("1", Uint32)
	assert.NoError(t, err)
	assert.Equal(t, uint32(1), val)

	_, err = FromString("-1", Uint32)
	assert.Error(t, err)

	val, err = FromString("1", Uint64)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), val)

	_, err = FromString("-1", Uint64)
	assert.Error(t, err)

	// float

	val, err = FromString("3.14", Float32)
	assert.NoError(t, err)
	assert.Equal(t, float32(3.14), val)

	_, err = FromString("str", Float32)
	assert.Error(t, err)

	val, err = FromString("3.14", Float64)
	assert.NoError(t, err)
	assert.Equal(t, 3.14, val)

	_, err = FromString("str", Float64)
	assert.Error(t, err)

	// bool

	val, err = FromString("true", Bool)
	assert.NoError(t, err)
	assert.Equal(t, true, val)

	_, err = FromString("1", Bool)
	assert.NoError(t, err)
	assert.Equal(t, true, val)

	val, err = FromString("false", Bool)
	assert.NoError(t, err)
	assert.Equal(t, false, val)

	_, err = FromString("0", Bool)
	assert.NoError(t, err)
	assert.Equal(t, false, val)

	_, err = FromString("invalid", Bool)
	assert.Error(t, err)

	// else

	_, err = FromString("0", "invalid")
	assert.Error(t, err, "cast: type %v is not supported", "invalid")
}
