package main

import "fmt"

func main() {
	var name string = "Jhon Smit"
	fmt.Println(name)
	name = "Jhon Lance"
	fmt.Println(name)

	var name1 = "Jhon Smit"
	fmt.Println(name1)
	name1 = "Jhon Lance"
	fmt.Println(name1)

	name2 := "Jhon Smit"
	fmt.Println(name2)
	name2 = "Jhon Lance"
	fmt.Println(name2)

	var (
		firstName  = "Jhon"
		middleName = "Smith"
		lastName   = "Lance"
	)
	fmt.Print(firstName, " "+middleName, " "+lastName)
}
