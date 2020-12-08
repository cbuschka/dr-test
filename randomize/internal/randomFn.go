package internal

import (
	"fmt"
	"math/rand"
	"reflect"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(maxLength int) string {
	b := make([]byte, 1+rand.Intn(maxLength))
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func randomInt() int64 {
	return rand.Int63()
}

func randomIntCapped(max int) int {
	return rand.Intn(max)
}

func randomFloat() float64 {
	return rand.Float64()
}

func randomBool() bool {
	return rand.Float32() <= 0.5
}

func randomSlice(sliceType reflect.Type, size int) reflect.Value {
	sliceOfType := reflect.SliceOf(sliceType)
	fmt.Printf("Slice type %v", sliceOfType)
	slice := reflect.MakeSlice(sliceOfType, 0, 5)
	fmt.Printf("Created slice %v", slice)
	for i := 0; i < size; i++ {
		slice = reflect.Append(slice, reflect.ValueOf(randomSimpleValue(sliceType)))
	}
	return slice
}

func randomSimpleValue(typ reflect.Type) interface{} {
	switch typ.Kind() {
	case reflect.String:
		return randomString(10)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return randomInt()
	case reflect.Float32, reflect.Float64:
		return randomFloat()
	case reflect.Bool:
		return randomBool()
	default:

		nestedStruct := reflect.New(typ)
		fmt.Printf("Nested struct\n%+v\n", nestedStruct)
		Randomize(nestedStruct.Interface(), Configuration{
			MaxListSize:     3,
			MaxStringLength: 10,
		})

		return nestedStruct.Elem().Interface()

	}
}
