package internal

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestRandomizeSimpleValues(t *testing.T) {

	t.Run("randomize string", func(t *testing.T) {
		type StringStruct struct {
			String string
		}
		got := randomizeWithDefaults(&StringStruct{})
		assert.NotNil(t, got)
		stringField := getField(got, "String").String()
		assert.True(t, len(stringField) > 0, "string is empty")
	})

	t.Run("randomize integer", func(t *testing.T) {
		type IntStruct struct {
			Int int64
		}
		in := &IntStruct{Int: 0}
		got := randomizeWithDefaults(in)
		assert.NotNil(t, got)
		intVal := getField(got, "Int").Int()
		assert.True(t, intVal != 0, "int has not changed")
	})

	t.Run("randomize float", func(t *testing.T) {
		type FloatStruct struct {
			Float float64
		}
		in := &FloatStruct{Float: 0.1}
		got := randomizeWithDefaults(in)
		assert.NotNil(t, got)
		floatVal := getField(got, "Float").Float()
		assert.True(t, floatVal != 0.1, "float has not changed")

	})

	t.Run("randomize boolean", func(t *testing.T) {
		type BooleanStruct struct {
			Boolean bool
		}
		got := randomizeWithDefaults(&BooleanStruct{})
		assert.NotNil(t, got)
	})

	t.Run("randomize multiple simple fields", func(t *testing.T) {
		type MultiStruct struct {
			Boolean bool
			Int32   int32
			Float64 float64
			String  string
		}
		got := randomizeWithDefaults(&MultiStruct{true, 0, 0.0, ""})
		assert.NotNil(t, got)
		assert.True(t, getField(got, "Int32").Int() > 0)
		assert.True(t, getField(got, "Float64").Float() > 0.0)
		assert.True(t, len(getField(got, "String").String()) > 0)
	})

}

func TestRandomizeSlices(t *testing.T) {
	type StructWithSlice struct {
		String       string
		Slice        []string
		SliceInt64   []int64
		SliceFloat64 []float64
		SliceBool    []bool
	}

	got := randomizeWithDefaults(&StructWithSlice{})
	assert.NotNil(t, got)
	assert.True(t, len(getField(got, "String").String()) > 0)
	assert.NotEmpty(t, getField(got, "Slice").Slice(0, 1))
	assert.NotEmpty(t, getField(got, "SliceInt64").Slice(0, 1))
	assert.NotEmpty(t, getField(got, "SliceFloat64").Slice(0, 1))
	assert.NotEmpty(t, getField(got, "SliceBool").Slice(0, 1))
}

func TestNestedStructs(t *testing.T) {

	type InnerStruct struct {
		Int       int64
		InnerName string
	}

	type MiddleStruct struct {
		MiddleName  string
		InnerStruct InnerStruct
	}

	type OuterStruct struct {
		OuterName    string
		MiddleStruct MiddleStruct
	}

	got := randomizeWithDefaults(&OuterStruct{})
	assert.NotNil(t, got)
	assert.True(t, len(getField(got, "OuterName").String()) > 0)
	assert.True(t, len(getField(got, "MiddleStruct").FieldByName("MiddleName").String()) > 0)
	assert.True(t, len(getField(got, "MiddleStruct").FieldByName("InnerStruct").FieldByName("InnerName").String()) > 0)

}

func getField(strukt interface{}, name string) reflect.Value {
	elem := reflect.ValueOf(strukt).Elem()
	return elem.FieldByName(name)
}

func randomizeWithDefaults(strukt interface{}) interface{} {
	return Randomize(strukt, Configuration{
		MaxListSize:     5,
		MaxStringLength: 5,
	})
}
