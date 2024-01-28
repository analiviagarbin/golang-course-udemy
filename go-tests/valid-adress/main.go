package main

import (
	"fmt"
	"valid-adress/adress"
)

func main() {
	typeAdress := adress.AdressType("Rua 15")
	fmt.Println(typeAdress)
}