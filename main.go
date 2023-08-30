package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type Member struct {
	ID         string `json:"id"`
	Firstname  string `json:"firstname"`
	Middlename string `json:"middlename"`
	Lastname   string `json:"lastname"`
	Email      string `json:"email"`
	Code       string `json:"code"`
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

func addMember(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	var params map[string]string
	json.NewDecoder(r.Body).Decode(&params)
	members = append(members, Member{ID: strconv.Itoa(rand.Intn(100)), Firstname: params["firstname"], Middlename: params["middlename"], Lastname: params["lastname"], Email: params["email"], Code: params["code"], Age: params["age"], Interest: params["interest"]})
	json.NewEncoder(w).Encode(&Member{})
}

func deleteMember(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range members {
		if item.ID == params["id"] {
			members = append(members[:i], members[i+1:]...)
			return
		}
	}
	json.NewEncoder(w).Encode(&Member{})
}

func updateMember(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	var params map[string]string
	json.NewDecoder(r.Body).Decode(&params)
	for i, item := range members {
		if item.ID == params["id"] {
			members = append(members[:i], members[i+1:]...)
			members = append(members, Member{ID: params["id"], Firstname: params["firstname"], Middlename: params["middlename"], Lastname: params["lastname"], Email: params["email"], Code: params["code"], Age: params["age"], Interest: params["interest"]})
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
	for _, item := range users {
		if item.Username == params["username"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

func main() {
	r := mux.NewRouter()
	// added Users
	users = append(users, User{ID: "1", Username: "assembly", Role: "admin"})
	users = append(users, User{ID: "2", Username: "client", Role: "client"})
	users = append(users, User{ID: "3", Username: "r.mordvinov", Role: "admin"})
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/user/{id}", getUser).Methods("GET")
	// added Members
	members = append(members, Member{ID: "1", Firstname: "Ivan", Middlename: "Ivanovich", Lastname: "Ivanov", Email: "i.ivanov@hexa.com", Code: "4321", Age: "22", Interest: "grappling"})
	members = append(members, Member{ID: "2", Firstname: "Oleg", Middlename: "Olegovich", Lastname: "Olegov", Email: "o.olegov@hexa.com", Code: "1234", Age: "37", Interest: "box"})
	members = append(members, Member{ID: "3", Firstname: "Marta", Middlename: "Therr", Lastname: "Maitz", Email: "m.maitz@hexa.com", Code: "3322", Age: "18", Interest: "stretching"})
	r.HandleFunc("/members", getMembers).Methods("GET")
	r.HandleFunc("/member/{id}", getMember).Methods("GET")
	r.HandleFunc("/member/{id}", updateMember).Methods("POST")
	r.HandleFunc("/member", addMember).Methods("POST")
	r.HandleFunc("/member/delete/{id}", deleteMember).Methods("GET")

	r.HandleFunc("/auth", login).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}
