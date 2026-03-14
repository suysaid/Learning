package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
}

func (s Student) display() {
	fmt.Println("Name:", s.Name)
	fmt.Println("Age:", s.Age)
	fmt.Println("Score:", s.Score)
	fmt.Println()
}

func main() {
	var students []Student

	for i := 1; i <= 3; i++ {
		var s Student
		fmt.Print("Enter Name: ")
		fmt.Scan(&s.Name)

		fmt.Print("Enter Age: ")
		fmt.Scan(&s.Age)

		fmt.Print("Enter Score: ")
		fmt.Scan(&s.Score)

		students = append(students, s)
	}
	for _, s := range students {
		s.display()
	}

}
