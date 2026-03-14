package main

import "fmt"

func checkNumber(n int) {
	if n%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}
}

func main() {
	var number int

	for i := 1; i <= 5; i++ {
		fmt.Print("Enter Number: ")
		fmt.Scan(&number)
		checkNumber(number)
	}

}
