package main

import (
	"fmt"
	"strconv"
)

func main() {
	var team = []string{"Rijal", "Clara", "Adit", "Ivan", "Yudha", "Eka", "Fathur", "Sigit", "Giva", "Inas"}

	for i, v := range team {
		fmt.Print(strconv.Itoa(i + 1))
		fmt.Println(". Hai " + v)
	}
}
