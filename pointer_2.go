package main

import "fmt"

type Student struct {
	Name  string
	Score int
}

func scoreupdate(s *Student, newScore int) {
	s.Score = newScore
}

func main() {
	student := Student{Name: "Alice", Score: 70}
	fmt.Println("Before:", student.Score)

	scoreupdate(&student, 95) // &student gives the memory address
	fmt.Println("After:", student.Score)
}
