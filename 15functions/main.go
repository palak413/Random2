package main

import "fmt"

func main() {
	fmt.Println("Welcome to function in golang")
	greet()
	// greeterTwo()

	result := adder(3, 4)

	fmt.Println("Result is:", result)
	proRes, _ := proAdder(2, 4, 6, 8)
	fmt.Println("Pro Result is", proRes)

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "Hi Pro result function"
}

// func greeterTwo(){
// 	fmt.Println("another method")

// }

func greet() {
	fmt.Println("Helooooooo")
}
