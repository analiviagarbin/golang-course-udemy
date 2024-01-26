// unit test
package adress

import "testing"

func TestAdressType(t *testing.T) {
	adressTesting := "Avenida Paulista"
	typeExpected := "Avenida"
	typeFunc := AdressType(adressTesting) 
	
	if typeFunc != typeExpected {
		t.Error("Invalid response!") 
	}
}