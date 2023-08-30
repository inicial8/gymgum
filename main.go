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

type Member struct {
	ID         string `json:"id"`
	Firstname  string `json:"firstname"`
	Middlename string `json:"middlename"`
	Lastname   string `json:"lastname"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Age        string `json:"age"`
	Interest   string `json:"interest"`
}

var users []User
var members []Member

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func getMembers(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(members)
}

func getMember(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range members {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Member{})
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
	// added Users
	users = append(users, User{ID: "1", Name: "assambly", Role: "admin"})
	users = append(users, User{ID: "2", Name: "client", Role: "client"})
	users = append(users, User{ID: "3", Name: "r.mordvinov", Role: "admin"})
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/user/{id}", getUser).Methods("GET")
	// added Members
	members = append(members, Member{ID: "1", Firstname: "Ivan", Middlename: "Ivanovich", Lastname: "Ivanov", Email: "i.ivanov@hexa.com", Password: "4321", Age: "22", Interest: "grappling"})
	members = append(members, Member{ID: "2", Firstname: "Oleg", Middlename: "Olegovich", Lastname: "Olegov", Email: "o.olegov@hexa.com", Password: "1234", Age: "37", Interest: "box"})
	members = append(members, Member{ID: "3", Firstname: "Marta", Middlename: "Therr", Lastname: "Maitz", Email: "m.maitz@hexa.com", Password: "3322", Age: "18", Interest: "stretching"})
	r.HandleFunc("/members", getMembers).Methods("GET")
	r.HandleFunc("/member/{id}", getMember).Methods("GET")

	r.HandleFunc("/auth", login).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}
