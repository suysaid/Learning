package main

import "fmt"

func changeValue(num *int) {
	*num = 100 // *num changes the value at the memory address
}

func main() {
	x := 10
	fmt.Println("Before:", x) // 10

	changeValue(&x)          // pass the address of x
	fmt.Println("After:", x) // 100
}
