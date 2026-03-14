package main

import "fmt"

type Students struct {
	Name  string
	Score int
}

func (s Students) display() {
	fmt.Println("Name:", s.Name)
	fmt.Println("Score:", s.Score)
	fmt.Println()
}

func main() {
	var students []Students

	for i := 1; i <= 3; i++ {
		var s Students
		fmt.Print("Enter Name: ")
		fmt.Scan(&s.Name)

		fmt.Print("Enter Score: ")
		fmt.Scan(&s.Score)

		students = append(students, s)
	}

	for _, s := range students {
		s.display()
	}
}
