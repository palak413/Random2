package main

import (
	"fmt"
)

func main() {
	fmt.Println("maps in golang")

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["rb"] = "ruby"

	fmt.Println("List of languages:", languages)

	fmt.Println("js shorts for", languages["js"])

	delete(languages, "rb")
	fmt.Println("List of languages:", languages)

	// loops is interseting in maps

	for key, value := range languages {
		fmt.Println("For key %v, value is %v\n", key, value)
	}
}
