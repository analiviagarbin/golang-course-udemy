package main

import (
	"fmt"
	"encoding/json"
	"log"
	"bytes"
)

type dog struct {
	Name string `json:"nome"`
	Race string `json:"race"`
	Age uint `json:"age"`
}

func main () {
	d := dog{"Maya", "German Shepherd", 3}
	fmt.Println(d)

	dogJSON, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dogJSON)
	fmt.Println(bytes.NewBuffer(dogJSON))

	d2 := map[string]string {
		"name": "Pity",
		"race": "SRD",
	}

	dog2JSON, err := json.Marshal(d2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytes.NewBuffer(dog2JSON))
} 