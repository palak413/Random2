package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	// fruitList[2] = "Banana"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is:", fruitList)
	fmt.Println("length of array is:", len(fruitList)) // now matter how many elements we have in array it will give the length of array

	var vegList = [3]string{"Potato", "Beans", "Cauliflower"}
	fmt.Println("Vegy list is:", vegList)

}
