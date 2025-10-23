package main

import "fmt"

// when a function excute it execute line by line but if there is a defer statement it will excecute at last just before the next line of code after the function is executed
func main() {
	defer fmt.Println("deferred")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("defer in golang")
	myDefer()
}

//  world, one, two
// 0,1,2,3,4
// hello, 43210, two, one, world
	func myDefer() {
		 for i:=0; i<5 ; i++ {
			defer fmt.Println(i)

		 }

	}


