package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // implicit import
)

// driver mysql: go get github.com/go-sql-driver/mysql

func main() {
	// stringConnection := "user:password@/bd"
	stringConnection := "golang:golang@/goCourse?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", stringConnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil { // test connection
		log.Fatal(err)
	}

	fmt.Println("The connection is open.")

	lines, err := db.Query("select * from users")
	if err != nil {
		log.Fatal(err)
	}
	defer lines.Close()

}
