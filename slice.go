package main

import "fmt"

func main() {
    var numbers []int

    for i := 1; i <= 5; i++ {
        var number int
        fmt.Print("Enter Number: ")
        fmt.Scan(&number)
        numbers = append(numbers, number)
    }
	fmt.Println("Numbers entered:")

    for _, value := range numbers {
        fmt.Println(value)
    }
}