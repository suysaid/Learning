package main

import "fmt"

func main() {
	var age int
	fmt.Print("enter your age")
	fmt.Scan(&age)

	if age < 18 {
		fmt.Println("you are a minor")
	} else {
		fmt.Println("you are an adult")
	}
}