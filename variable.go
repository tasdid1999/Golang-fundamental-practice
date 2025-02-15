package main

import "fmt"

func main() {

	var firstName string

	firstName = "tasdid"

	fmt.Printf("my fisrt name is %s\n", firstName)

	lastName := "alam"

	fmt.Printf("my last name is %s\n", lastName)

	var age int

	age = 24

	fmt.Printf("my age is %d\n", age)

	var isMarried bool

	isMarried = false

	if isMarried {
		fmt.Println("he is unHappy!")
	} else {
		fmt.Println("happy")
	}
}
