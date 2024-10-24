package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var users = make(map[string]User)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}).Methods("GET")

	r.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		user, exists := users[id]
		if !exists {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(user)
	}).Methods("GET")

	r.HandleFunc(("/users"), func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(users)
	}).Methods("GET")

	r.HandleFunc("/users/criar", func(w http.ResponseWriter, r *http.Request) {
		var user User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		users[user.ID] = user
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}).Methods("POST")

	r.HandleFunc("/users/delete/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		if _, exists := users[id]; !exists {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		delete(users, id)
		w.WriteHeader(http.StatusNoContent)
	}).Methods("DELETE")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
