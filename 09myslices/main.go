package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices in golang")

	var fruitList = []string{"Apple", "Tomato", "Peach"} // this is how we create slice we don't need to maintain size unlike array
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana") // we can add more elements in slice using append functions
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3]) // this will remove first element and keep 2nd and 3rd element
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 456
	highScores[2] = 678
	highScores[3] = 867
	// highScores[4] = 1000 // this will give error because we have defined size of slice as 4

	highScores = append(highScores, 555, 666, 777) // this is how we can add more elements in slice

	// fmt.Println(highScores)
	sort.Ints(highScores)
	// fmt.Println(highScores)

	// fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove a value from slice based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
