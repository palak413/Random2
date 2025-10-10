package main

import "fmt"

const LoginToken string = "ghabhhed.jd" // public denoted by first letter captial

func main() {
	var username string = "Palak"
	fmt.Println(username)
	fmt.Printf("Type of username is %T\n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Type of isLoggedIn is %T\n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Type of smallVal is %T\n", smallVal)

	var smallFloat float32 = 255.4555555555
	fmt.Println(smallFloat)
	fmt.Printf("Type of smallFloat is %T\n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Type of anotherVariable is %T\n", anotherVariable)

	var anotherFloat float64
	fmt.Println(anotherFloat)
	fmt.Printf("Type of anotherFloat is %T\n", anotherFloat)

	// no var style

	numberOfUser := 30000
	fmt.Println(numberOfUser)
	fmt.Printf("Type of numberOfUser is %T\n", numberOfUser)

	fmt.Println("Login Token is:", LoginToken)

}
