package cast_test

import (
	"github.com/golobby/cast"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromString(t *testing.T) {
	message := "cast: cannot cast `%v` to type `%v`"

	// string

	val, err := cast.FromString("string", cast.String)
	assert.NoError(t, err)
	assert.Equal(t, "string", val)

	// int

	val, err = cast.FromString("1", cast.Int)
	assert.NoError(t, err)
	assert.Equal(t, 1, val)

	_, err = cast.FromString("str", cast.Int)
	assert.Errorf(t, err, message, "str", cast.Int)

	val, err = cast.FromString("1", cast.Int8)
	assert.NoError(t, err)
	assert.Equal(t, int8(1), val)

	_, err = cast.FromString("str", cast.Int8)
	assert.Errorf(t, err, message, "str", cast.Int8)

	val, err = cast.FromString("1", cast.Int16)
	assert.NoError(t, err)
	assert.Equal(t, int16(1), val)

	_, err = cast.FromString("str", cast.Int16)
	assert.Errorf(t, err, message, "str", cast.Int16)

	val, err = cast.FromString("1", cast.Int32)
	assert.NoError(t, err)
	assert.Equal(t, int32(1), val)

	_, err = cast.FromString("str", cast.Int32)
	assert.Errorf(t, err, message, "str", cast.Int32)

	val, err = cast.FromString("1", cast.Int64)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), val)

	_, err = cast.FromString("str", cast.Int64)
	assert.Errorf(t, err, message, "str", cast.Int64)

	// uint

	val, err = cast.FromString("1", cast.Uint)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), val)

	_, err = cast.FromString("-1", cast.Uint)
	assert.Errorf(t, err, message, "-1", cast.Uint)

	val, err = cast.FromString("1", cast.Uint8)
	assert.NoError(t, err)
	assert.Equal(t, uint8(1), val)

	_, err = cast.FromString("-1", cast.Uint8)
	assert.Errorf(t, err, message, "-1", cast.Uint8)

	val, err = cast.FromString("1", cast.Uint16)
	assert.NoError(t, err)
	assert.Equal(t, uint16(1), val)

	_, err = cast.FromString("-1", cast.Uint16)
	assert.Errorf(t, err, message, "-1", cast.Uint16)

	val, err = cast.FromString("1", cast.Uint32)
	assert.NoError(t, err)
	assert.Equal(t, uint32(1), val)

	_, err = cast.FromString("-1", cast.Uint32)
	assert.Errorf(t, err, message, "-1", cast.Uint32)

	val, err = cast.FromString("1", cast.Uint64)
	assert.NoError(t, err)
	assert.Equal(t, uint64(1), val)

	_, err = cast.FromString("-1", cast.Uint64)
	assert.Errorf(t, err, message, "-1", cast.Uint64)

	// float

	val, err = cast.FromString("3.14", cast.Float32)
	assert.NoError(t, err)
	assert.Equal(t, float32(3.14), val)

	_, err = cast.FromString("str", cast.Float32)
	assert.Errorf(t, err, message, "str", cast.Float32)

	val, err = cast.FromString("3.14", cast.Float64)
	assert.NoError(t, err)
	assert.Equal(t, 3.14, val)

	_, err = cast.FromString("str", cast.Float64)
	assert.Errorf(t, err, message, "str", cast.Float64)

	// bool

	val, err = cast.FromString("true", cast.Bool)
	assert.NoError(t, err)
	assert.Equal(t, true, val)

	_, err = cast.FromString("1", cast.Bool)
	assert.NoError(t, err)
	assert.Equal(t, true, val)

	val, err = cast.FromString("false", cast.Bool)
	assert.NoError(t, err)
	assert.Equal(t, false, val)

	_, err = cast.FromString("0", cast.Bool)
	assert.NoError(t, err)
	assert.Equal(t, false, val)

	_, err = cast.FromString("invalid", cast.Bool)
	assert.Errorf(t, err, message, "invalid", cast.Bool)

	// array

	val, err = cast.FromString("4,-5,6", cast.IntArray)
	assert.NoError(t, err)
	assert.Equal(t, []int{4, -5, 6}, val)

	val, err = cast.FromString("a,b,c", cast.IntArray)
	assert.Error(t, err)

	val, err = cast.FromString("4,-5,6", cast.Int8Array)
	assert.NoError(t, err)
	assert.Equal(t, []int8{4, -5, 6}, val)

	val, err = cast.FromString("a,b,c", cast.Int8Array)
	assert.Error(t, err)

	val, err = cast.FromString("4,-5,6", cast.Int16Array)
	assert.NoError(t, err)
	assert.Equal(t, []int16{4, -5, 6}, val)

	val, err = cast.FromString("a,b,c", cast.Int16Array)
	assert.Error(t, err)

	val, err = cast.FromString("4,-5,6", cast.Int32Array)
	assert.NoError(t, err)
	assert.Equal(t, []int32{4, -5, 6}, val)

	val, err = cast.FromString("a,b,c", cast.Int32Array)
	assert.Error(t, err)

	val, err = cast.FromString("4,-5,6", cast.Int64Array)
	assert.NoError(t, err)
	assert.Equal(t, []int64{4, -5, 6}, val)

	val, err = cast.FromString("a,b,c", cast.Int64Array)
	assert.Error(t, err)

	val, err = cast.FromString("4,5,6", cast.UintArray)
	assert.NoError(t, err)
	assert.Equal(t, []uint{4, 5, 6}, val)

	val, err = cast.FromString("a,b,c", cast.UintArray)
	assert.Error(t, err)

	val, err = cast.FromString("4,5,6", cast.Uint8Array)
	assert.NoError(t, err)
	assert.Equal(t, []uint8{4, 5, 6}, val)

	val, err = cast.FromString("a,b,c", cast.Uint8Array)
	assert.Error(t, err)

	val, err = cast.FromString("4,5,6", cast.Uint16Array)
	assert.NoError(t, err)
	assert.Equal(t, []uint16{4, 5, 6}, val)

	val, err = cast.FromString("a,b,c", cast.Uint16Array)
	assert.Error(t, err)

	val, err = cast.FromString("4,5,6", cast.Uint32Array)
	assert.NoError(t, err)
	assert.Equal(t, []uint32{4, 5, 6}, val)

	val, err = cast.FromString("a,b,c", cast.Uint32Array)
	assert.Error(t, err)

	val, err = cast.FromString("4,5,6", cast.Uint64Array)
	assert.NoError(t, err)
	assert.Equal(t, []uint64{4, 5, 6}, val)

	val, err = cast.FromString("a,b,c", cast.Uint64Array)
	assert.Error(t, err)

	val, err = cast.FromString("3.14,9.8", cast.Float32Array)
	assert.NoError(t, err)
	assert.Equal(t, []float32{3.14, 9.8}, val)

	val, err = cast.FromString("3.14,9.8", cast.Float64Array)
	assert.NoError(t, err)
	assert.Equal(t, []float64{3.14, 9.8}, val)

	val, err = cast.FromString("true,false,0,1", cast.BoolArray)
	assert.NoError(t, err)
	assert.Equal(t, []bool{true, false, false, true}, val)

	val, err = cast.FromString("a,b,c", cast.StringArray)
	assert.NoError(t, err)
	assert.Equal(t, []string{"a", "b", "c"}, val)

	// else

	_, err = cast.FromString("0", "invalid")
	assert.Error(t, err, "cast: type %v is not supported", "invalid")

	_, err = cast.FromString("0,1", "[]invalid")
	assert.Error(t, err, "cast: type %v is not supported", "[]invalid")
}
