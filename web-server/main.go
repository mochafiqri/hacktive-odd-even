package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sync"
)

const (
	PORT = ":8080"
)

func main() {
	r := gin.Default()

	var user = NewUserContract()
	r.POST("/register", user.Register)
	r.GET("/user", user.GetUsers)
	r.GET("/user/:number", user.GetUser)

	r.Run(PORT)
}

type UserModel struct {
	Name string `json:"name"`
}

type User interface {
	Register(c *gin.Context)
	GetUser(c *gin.Context)
	GetUsers(c *gin.Context)
}

type UserContract struct {
	Users []UserModel
}

func (u *UserContract) Register(c *gin.Context) {
	var req = UserModel{}

	var err = c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	u.Users = append(u.Users, req)

	c.JSON(http.StatusOK, req)
}

func (u *UserContract) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, u.Users)
}
func (u *UserContract) GetUser(c *gin.Context) {
	var tmp = c.Param("number")
	var number, err = strconv.Atoi(tmp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	if len(u.Users) < number {
		c.JSON(http.StatusOK, map[string]interface{}{
			"error": "data not found",
		})
		return
	}

	c.JSON(http.StatusOK, u.Users[number-1])
}

func NewUserContract() User {
	return &UserContract{}
}

func print(wg *sync.WaitGroup, user *UserModel) {
	fmt.Println(user.Name)
	wg.Done()
}
