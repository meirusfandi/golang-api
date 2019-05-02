package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//create user struct/ model
type User struct {
	ID       string  `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Fullname string  `json:"fullname"`
	Email    string  `json:"email"`
	Alamat   *Alamat `json:"alamat"`
}

//Alamat struct / model
type Alamat struct {
	IDAlamat  string `json:"idalamat"`
	Jalan     string `json:"jalan"`
	RT        string `json:"rt"`
	RW        string `json:"rw"`
	Kelurahan string `json:"kelurahan"`
	Kecamatan string `json:"kecamatan"`
	Kabupaten string `json:"kabupaten"`
	Provinsi  string `json:"provinsi"`
	KodePos   string `json:"kodepos"`
}

//all user struct model as slice
var AllUsers []User

//get all user methods
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(AllUsers)
}

//get user by id methods
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	//get params from request
	params := mux.Vars(r)

	//looping all find params
	for _, item := range AllUsers {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&User{})
}

//create new user methods
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = strconv.Itoa(rand.Intn(100000))
	AllUsers = append(AllUsers, user)
	json.NewEncoder(w).Encode(user)
}

//update user methods
func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for index, item := range AllUsers {
		if item.ID == params["id"] {
			AllUsers = append(AllUsers[:index], AllUsers[index+1:]...)
			var user User
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.ID = item.ID
			AllUsers = append(AllUsers, user)
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode(AllUsers)
}

//delete user methods
func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for index, item := range AllUsers {
		if item.ID == params["id"] {
			AllUsers = append(AllUsers[:index], AllUsers[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(AllUsers)
}

func main() {
	//init new Router
	r := mux.NewRouter()

	//create Dummy data
	AllUsers = append(AllUsers, User{
		ID:       "1",
		Username: "meirf",
		Password: "coba123",
		Fullname: "Mei Rusfandi",
		Email:    "meirusfandi100@gmail.com",
		Alamat: &Alamat{
			IDAlamat:  "1",
			Jalan:     "Jl. Mangga 1",
			RT:        "02",
			RW:        "07",
			Kelurahan: "Jelupang",
			Kecamatan: "Serpong Utara",
			Kabupaten: "Tangerang Selatan",
			Provinsi:  "Banten",
			KodePos:   "15323",
		},
	})

	AllUsers = append(AllUsers, User{
		ID:       "2",
		Username: "mrcondong",
		Password: "coba12345",
		Fullname: "Mr Condong",
		Email:    "mrcondong@gmail.com",
		Alamat: &Alamat{
			IDAlamat:  "1",
			Jalan:     "Jl. Mangga 1",
			RT:        "02",
			RW:        "07",
			Kelurahan: "Jelupang",
			Kecamatan: "Serpong Utara",
			Kabupaten: "Tangerang Selatan",
			Provinsi:  "Banten",
			KodePos:   "15323",
		},
	})

	//create new Handler or Endpoints
	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/api/users", createUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/api/users/{id}", deleteUser).Methods("DELETE")

	fmt.Println("server start at localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
