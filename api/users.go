package main

import (
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	db, err := dbConnect()
	if err != nil {
		http.Error(w, "Error connecting to database: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT Id, Name, Lastname, Username, Password FROM users")
	if err != nil {
		http.Error(w, "Error fetching users: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Lastname, &user.Username, &user.Password); err != nil {
			http.Error(w, "Error scanning user: "+err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	db, err := dbConnect()
	if err != nil {
		http.Error(w, "Error connecting to database: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Error decoding request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO users (Name, Lastname, Username, Password) VALUES (?, ?, ?, ?)",
		user.Name, user.Lastname, user.Username, hashPassword(user.Password))
	if err != nil {
		http.Error(w, "Error creating user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	db, err := dbConnect()
	if err != nil {
		http.Error(w, "Error connecting to database: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Error decoding request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE users SET Name=?, Lastname=?, Username=?, Password=? WHERE Id=?",
		user.Name, user.Lastname, user.Username, hashPassword(user.Password), user.ID)
	if err != nil {
		http.Error(w, "Error updating user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	db, err := dbConnect()
	if err != nil {
		http.Error(w, "Error connecting to database: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Error decoding request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("DELETE FROM users WHERE Id=?", user.ID)
	if err != nil {
		http.Error(w, "Error deleting user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
