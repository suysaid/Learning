package main

import "fmt"

func main() {

	numbers := []int{10, 20, 30, 40, 50}

	fmt.Println(numbers[:])
	fmt.Println(numbers[:3])
	fmt.Println(numbers[1:4])
	fmt.Println(numbers[2:])

}
