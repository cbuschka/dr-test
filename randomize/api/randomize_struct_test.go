package api

import (
	"drtest/randomize/internal"
	"encoding/json"
	"fmt"
	"testing"
)

func TestRandomize(t *testing.T) {

	randomized := randomize(&Person{}, internal.Configuration{
		MaxListSize:     4,
		MaxStringLength: 5,
	})
	person := randomized.(*Person)
	fmt.Printf("---------\n%v", person)

}

type Pet struct {
	Name string
	Age  int64
}

type Coordinates struct {
	Lat float64
	Lon float64
}

type Person struct {
	FirstName    string
	LastName     string
	Hobbies      []string
	LuckyNumbers []int64
	Cool         bool
	Balance      float64
	Coordinates  Coordinates
	Pets         []Pet
	BFF          []Person
}

func (p Person) String() string {
	indent, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		return err.Error()
	}
	return string(indent)
}
