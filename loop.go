package main

import "fmt"

func main() {
	var choice int

	fmt.Print("How many hello")
	fmt.Scan(&choice)

	for i := 1; i <= choice; i++ {
		fmt.Println("Hello")
	}
}