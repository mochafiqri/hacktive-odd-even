package main

import "fmt"

type Member struct {
	Name string
}

func main() {
	var team = []string{"Rijal", "Clara", "Adit", "Ivan", "Yudha", "Eka", "Fathur", "Sigit", "Giva", "Inas"}

	var members []*Member
	var addMember = func(name string) {
		members = append(members, &Member{Name: name})
	}
	for _, v := range team {
		addMember(v)
	}

	for _, v := range members {
		fmt.Println("Hai : ", v.Name)
	}
}
