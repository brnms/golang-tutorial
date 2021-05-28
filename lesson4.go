package main

import "fmt"

func main() {
	age := 29
	name := "bruno"

	// print
	fmt.Print("hello, \n")
	fmt.Print("world \n")
	fmt.Print("new line \n")

	// Println
	fmt.Println("hello ninjas!")
	fmt.Println("goodbye ninjas!")
	fmt.Println("my age is", age, "and my name is", name)

	// Printf (formatted strings) &_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("you scored %f points \n", 225.55)
	fmt.Printf("you scored %0.1f points \n", 225.55)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v", age, name)
	fmt.Println("the saved string is: ", str)

}
