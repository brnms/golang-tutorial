package main

import "fmt"

func main() {
	age := 25

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is not less than 45")
	}

	names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("\ncontinuing at pos", index)
			continue
		}

		if index > 2 {
			fmt.Println("\nbreaking at pos", index)
			break
		}

		fmt.Printf("\nthe value at pos %v is %v", index, value)
	}
}
