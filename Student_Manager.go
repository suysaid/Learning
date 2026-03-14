package main

import "fmt"

type Student struct {
	Name  string
	Score int
}

func main() {
	var students []Student

	for {
		fmt.Println("1. Add Student")
		fmt.Println("2. Show Students")
		fmt.Println("3. Exit")
		fmt.Println("4. Clear Students")

		var choice int
		fmt.Print("Choose: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			var s Student

			fmt.Print("Enter Name: ")
			fmt.Scan(&s.Name)

			fmt.Print("Enter Score: ")
			fmt.Scan(&s.Score)

			students = append(students, s)

		case 2:
			if len(students) == 0 {
				fmt.Println("No students found.")
			} else {
				for _, s := range students {
					fmt.Println("Name:", s.Name, "Score:", s.Score)
				}
			}

		case 3:
			fmt.Println("Goodbye!")
			return

		case 4:
			students = nil
			fmt.Println("All Studets Cleared!")

		default:
			fmt.Println("Invalid Choice")

		}
	}

}
