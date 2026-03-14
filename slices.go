package main

import "fmt"

func main() {
	var numbers []int

	for i := 0; i < 5; i++ {
		var number int
		fmt.Print("Enter Number: ")
		fmt.Scan(&number)

		numbers = append(numbers, number)
	}

	fmt.Println("Slice:", numbers)
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))
}
