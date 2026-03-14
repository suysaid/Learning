package main

import "fmt"

func main() {

	numbers := []int{5, 10, 15, 20, 25, 30}

	fmt.Println(numbers[1:4])
	fmt.Println(numbers[:2])
	fmt.Println(numbers[3:])

}
