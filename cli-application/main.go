package main

import (
	"app-cli/app"
	"log"
	"os"
)

func main() {
	application := app.Create()
	erro := application.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}
}
