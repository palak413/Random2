package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golang")

	rand.Seed(time.Now().UnixNano()) // by default it is coming from math package and w'll use that only w'll not use any other rand package
	diceNumber := rand.Intn(6) + 1   // it will give number between 0 to 5 so we add 1 to make it 1 to 6
	fmt.Println("Dice number is :", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spot")
		fallthrough // it will execute next case also
	case 4:
		fmt.Println("You can move 4 spot")
	case 5:
		fmt.Println("You can move 5 spot")
	case 6:
		fmt.Println("You can move 6 spot and play again")
	default:
		fmt.Println("What was that?")

	}

}
