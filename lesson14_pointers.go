package main

import "fmt"

func updateName(x *string) {
	*x = "wedge"
}

func main() {
	// Group A types -> strings, ints, bools, floats, arrays, structs
	name := "tifa"

	//updateName(name)

	//fmt.Println("memory adress of name is: ", &name)

	m := &name
	// fmt.Println("memory adress:", m)
	// fmt.Println("value at memory adress:", *m)

	fmt.Println(name)
	updateName(m)
	fmt.Println(name)

}
