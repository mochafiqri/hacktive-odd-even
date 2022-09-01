package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

const (
	PORT = ":8080"
)

func main() {
	fmt.Println("TES")
	var user = NewUserContract()
	http.HandleFunc("/register", user.Register)
	http.HandleFunc("/user", user.GetUser)

	var err = http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

}

type UserModel struct {
	Name string `json:"name"`
}

type User interface {
	Register(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
}

type UserContract struct {
	Users []UserModel
}

func (u *UserContract) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		var req = UserModel{}

		decoder := json.NewDecoder(r.Body)

		var err = decoder.Decode(&req)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		u.Users = append(u.Users, req)
		json.NewEncoder(w).Encode(u.Users)
	}
}

func (u *UserContract) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(u.Users)
	}
}

func NewUserContract() User {
	return &UserContract{}
}

func print(wg *sync.WaitGroup, user *UserModel) {
	fmt.Println(user.Name)
	wg.Done()
}
