package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Jhon"
	names[1] = "Smith"
	names[2] = "lance"
	// names[3] = "Guis" //Erorr

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		10,
		20,
		30,
	}
	fmt.Println(values)

	var valuesDua = [...]int{
		40,
		50,
		60,
		70,
		80,
		90,
	}
	fmt.Println(valuesDua)
	fmt.Println(len(valuesDua))
	valuesDua[0] = 100
	fmt.Println(valuesDua[0])
}
