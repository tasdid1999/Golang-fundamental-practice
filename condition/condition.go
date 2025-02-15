package main

import "fmt"

func main() {

	mark := 70

	if mark >= 80 {

		fmt.Println("grade is A+")

	} else if mark >= 70 && mark < 80 {

		fmt.Println("grade is A")

	} else if mark >= 60 && mark < 70 {

		fmt.Println("grade is A-")

	} else {
		fmt.Println("fail")
	}
}
