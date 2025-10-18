package main

import "fmt"

func main() {
	fmt.Println("if else statements in golang")

	loginCount := 10

	var result string

	if loginCount < 10 { //u can't use these curly braces in next line
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "exactly 10 login count"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Number os less than 10")
	} else {
		fmt.Println("Number is greater than 10")
	}

}
