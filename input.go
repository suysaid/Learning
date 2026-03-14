package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Print("Enter Your Name")
	fmt.Scan(&name)

	fmt.Print("Enter Your Age")
	fmt.Scan(&age)

	fmt.Println("Hello", name)
	fmt.Println("You are", age, "years old")
}
