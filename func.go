package main

import "fmt"

func greetUser()  {
	var Name string
	fmt.Print("What's your name?")
	fmt.Scan(&Name)

	fmt.Println("Hello", Name)
}

func main()  {
	greetUser()
	greetUser()
	
}