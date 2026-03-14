package main

import "fmt"

func main() {
	scores := make(map[string]int)

	for i := 1; i <= 3; i++ {
		var name string
		var score int

		fmt.Print("Enter Name: ")
		fmt.Scan(&name)

		fmt.Print("Enter Score: ")
		fmt.Scan(&score)

		scores[name] = score
	}

	for name, score := range scores {
		fmt.Println(name, ":", score)
	}

}
