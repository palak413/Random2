package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file"

	file, err := os.Create("./mylcogofile.txt")

	if err != nil {
		panic(err) // panic will just stop the execution of the program  and will show u what the error u r facing
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("Length is: ", length)
	defer file.Close() // this will close the file after the main function is executed
	readFile("./mylcogofile.txt")

}

func readFile(filename string) {
	// _, err := ioutil.ReadFile(filename) // whenever u read a file it is not being read in form of string it is read in form of bytes
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("Text data inside the file is \n", string(databyte))

}
