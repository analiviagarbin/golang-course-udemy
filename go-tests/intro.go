package main

import {
	"fmt"
	"go-tests/adress"
}

func main() {
	typeAdress := adress.AdressType("Rua 15")
	fmt.Println(typeAdress)
}