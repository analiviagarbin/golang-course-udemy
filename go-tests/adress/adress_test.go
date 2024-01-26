// unit test
package adress

import "testing"

type testingScenario struct {
	adressFunc string,
	adressExpected string
}

func TestAdressType(t *testing.T) {

	testingScenarios := []testingScenario {
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Quadra 08", "Quadra"}
	}

	for _, typeExpected range testingScenarios {
		typeFunc := AdressType(typeExpected.adressFunc)
		if typeFunc != typeExpected.adressExpected {
			t.Error("Invalid response")
		}   
	}
}