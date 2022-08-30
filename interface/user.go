package main

import "fmt"

type UserModel struct {
	Name string
}

type User interface {
	Register(user *UserModel)
	GetUser() []*UserModel
}

type UserContract struct {
	Users []*UserModel
}

func (u *UserContract) Register(user *UserModel) {
	u.Users = append(u.Users, user)
}

func (u *UserContract) GetUser() []*UserModel {
	//TODO implement me
	return u.Users
}

func NewUserContract() User {
	return &UserContract{}
}

func main() {

	var u = NewUserContract()
	u.Register(&UserModel{Name: "Fiqri"})
	u.Register(&UserModel{Name: "Clara"})

	var users = u.GetUser()
	for _, v := range users {
		fmt.Println(v.Name)
	}

}
