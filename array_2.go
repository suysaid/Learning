package main

import (
	"fmt"
)

func main() {
	var score [5]int
	sum := 0

	for i := 0; i < 5; i++ {
		fmt.Print("Enter Number: ")
		fmt.Scan(&score[i])
	}

	fmt.Println("Array: ", score)

	for i := 0; i < 5; i++ {
		sum += score[i]
	}

	fmt.Println("Sum: ", sum)

	max := score[0]

	for i := 1; i < 5; i++ {
		if score[i] > max {
			max = score[i]
		}
	}

	fmt.Println("Highest Number: ", max)
}
