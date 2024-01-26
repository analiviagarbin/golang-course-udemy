package adress

import "strings"

// AdressType validates the entered address
func AdressType(adress string) string {
	validTypes := []string("rua", "avenida", "quadra")
	convertString := strings.ToLower(adress)

	fstWord := strings.Split(adress, " ")[0]
	adressHasValidType := false

	for _, type := range validTypes {
		if type == fstWord{
			adressHasValidType = true
		}
	}

	if adressHasValidType {
		return strings.Title(fstWord)
	} 

	return "Invalid Type"
}