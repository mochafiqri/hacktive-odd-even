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
	http.HandleFunc("/", greet)
	http.HandleFunc("/register", Register)

	var err = http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Word")
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		var req = UserModel{}
		var user = NewUserContract()

		decoder := json.NewDecoder(r.Body)

		var err = decoder.Decode(&req)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		user.Register(req)
		var users = user.GetUser()

		json.NewEncoder(w).Encode(users)
	}
}

type UserModel struct {
	Name string `json:"name"`
}

type User interface {
	Register(user UserModel)
	GetUser() []UserModel
}

type UserContract struct {
	Users []UserModel
}

func (u *UserContract) Register(user UserModel) {
	u.Users = append(u.Users, user)
}

func (u *UserContract) GetUser() []UserModel {
	return u.Users
}

func NewUserContract() User {
	return &UserContract{}
}

func print(wg *sync.WaitGroup, user *UserModel) {
	fmt.Println(user.Name)
	wg.Done()
}
