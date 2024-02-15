package server

import (
	"crud/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// NewUser creates a new user in db
func NewUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Failed to read the request body"))
		return
	}

	var user user

	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		w.Write([]byte("Error converting user to struct"))
		return
	}

	db, err := db.Connection()
	if err != nil {
		w.Write([]byte("Error connecting to the database"))
		return
	}
	defer db.Close()

	// insert into users (name, email) values ("name", "email")
	// prepare statement
	statement, err := db.Prepare("insert into users (name, email) values (?, ?)")
	if err != nil {
		w.Write([]byte("Error creating statement"))
		return
	}
	defer statement.Close()

	insertion, err := statement.Exec(user.Name, user.Email)
	if err != nil {
		w.Write([]byte("Error executing statement"))
		return
	}

	newId, err := insertion.LastInsertId()
	if err != nil {
		w.Write([]byte("Error obtaining inserted ID"))
		return
	}

	// status code
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User %d created succesfuly", newId)))
}

// SearchUsers list all users
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connection()
	if err != nil {
		w.Write([]byte("Error connecting to the database"))
		return
	}
	defer db.Close()

	// select * from users
	lin, err := db.Query("select * from users")
	if err != nil {
		w.Write([]byte("Error fetching all users"))
		return
	}
	defer lin.Close()

	var users []user
	for lin.Next() { // populating a slice of users
		var user user

		if err := lin.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Error scanning users"))
			return
		}

		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Error converting users to JSON"))
		return
	}
}

// SearchUser searches for a specific user
func SearchUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	ID, err := strconv.ParseUint(parameters["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Error converting parameter ID to integer"))
		return
	}

	db, err := db.Connection()
	if err != nil {
		w.Write([]byte("Error connecting to the database"))
		return
	}
	defer db.Close()

	lin, err := db.Query("select * from users where id = ?", ID)
	if err != nil {
		w.Write([]byte("Error searching for user"))
		return
	}

	var user user
	if lin.Next() {
		if err := lin.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Error scanning user"))
			return
		}
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Error converting user to JSON"))
		return
	}
}

// UpdateUser updates an existing user
func UpdateUser(w http.ResponseWriter, r *http.Request) {}

// DeleteUser deletes an existing user
func DeleteUser(w http.ResponseWriter, r *http.Request) {}
