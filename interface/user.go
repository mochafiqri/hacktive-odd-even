package main

import (
	"fmt"
	"sync"
)

type UserModel struct {
	Name string
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

func main() {

	var u = NewUserContract()
	wg := sync.WaitGroup{}
	wg.Add(4)
	go func() {
		u.Register(UserModel{Name: "Fiqri"})
		wg.Done()
	}()
	go func() {
		u.Register(UserModel{Name: "Clara"})
		wg.Done()

	}()
	go func() {
		u.Register(UserModel{Name: "Rijal"})
		wg.Done()

	}()
	go func() {
		u.Register(UserModel{Name: "Medy"})
		wg.Done()

	}()
	wg.Wait()

	var users = u.GetUser()
	wg.Add(len(users))

	for i := 0; i < len(users); i++ {
		go print(&wg, &users[i])
	}

	wg.Wait()
}

func print(wg *sync.WaitGroup, user *UserModel) {
	fmt.Println(user.Name)
	wg.Done()
}
