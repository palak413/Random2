package main

import "fmt"

func main() {
	fmt.Println("if else statements in golang")

	loginCount := 10

	var result string

	if loginCount < 10 {
		result = "Regular user"
	}

	fmt.Println(result)
}
