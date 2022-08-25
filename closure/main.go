package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Member struct {
	Name    string
	Address string
	Job     string
	Reason  string
}

func main() {

	index := os.Args[1:]
	var i = 0
	var err error
	if len(index) > 0 {
		i, err = strconv.Atoi(index[0])
		if err != nil {
			log.Fatalln(err)
		}
	}

	var team = []*Member{{
		Name:    "Fiqri",
		Address: "Bandung",
		Job:     "Backend",
		Reason:  "Keren",
	},
		{
			Name:    "Clara",
			Address: "Bekasi",
			Job:     "Frontend",
			Reason:  "Keren",
		},
		{
			Name:    "Rijal",
			Address: "Pinrang",
			Job:     "Haha Hihi",
			Reason:  "Keren",
		}}

	var print = func(i int) {
		i--
		if len(team) < i || i < 0 {
			fmt.Println("Team ga ada")
			return
		}
		fmt.Println("Name : ", team[i].Name)
		fmt.Println("Address : ", team[i].Name)
		fmt.Println("Job : ", team[i].Name)
		fmt.Println("Reason : ", team[i].Name)
	}

	print(i)

}
