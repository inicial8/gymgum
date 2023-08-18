package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

var users []User

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

func login(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	var params map[string]string
	json.NewDecoder(r.Body).Decode(&params)
	fmt.Println(params["username"])
	for _, item := range users {
		if item.Name == params["username"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

func main() {
	r := mux.NewRouter()
	users = append(users, User{ID: "1", Name: "assambly", Role: "admin"})
	users = append(users, User{ID: "2", Name: "r.mordvinov", Role: "admin"})
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/user/{id}", getUser).Methods("GET")
	r.HandleFunc("/auth", login).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}
