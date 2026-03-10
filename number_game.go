package main

import "fmt"

func checkNumber(n int) {
	if n % 2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}
}

func main() {
	var number int

	fmt.Print("Enter Number: ")
	fmt.Scan(&number)

	for i := 1; i <= number; i++ {
		fmt.Println(i)
	}

	checkNumber(number)
}