package main

import (
	"fmt"
	"strconv"
)

const Odd = "odd"
const Even = "even"

func main() {
	fmt.Println("Hello Word")
	for i := 0; i <= 10; i++ {
		fmt.Println(OddEven(i))
	}
}

func OddEven(number int) string {
	if number%2 == 0 {
		return strconv.Itoa(number) + " " + Even
	}
	return strconv.Itoa(number) + " " + Odd
}
