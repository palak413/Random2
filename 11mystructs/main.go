package main

import "fmt"

func main() {

	fmt.Println("structs in golang")

	// no inheritance in golang ; no super or parent key word

	palak := User{"Palak", "palak@go.dev", true, 19}
	fmt.Println(palak)
	fmt.Printf("palak details are: %+v\n", palak)
	fmt.Printf("Name is %v and email is %v\n", palak.Name, palak.Email)

}

	type User struct {
		Name   string
		Email  string
		Status bool
		Age    int
	}