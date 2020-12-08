package person

import "drtest/generated/person/avro"

func GeneratePerson(amount int) []interface{} {
	sliced := make([]interface{}, amount)
	for i := range sliced {
		sliced[i] = randomize(avro.NewPerson())
	}
	return sliced
}
func randomize(person interface{}) interface{} {
	return person
}