package main

import "fmt"

func getName() (firstName, middleName, lastName string) {
	firstName = "Jhon"
	middleName = "Smith"
	lastName = "Lance"

	return firstName, middleName, lastName

}
func main() {
	a, b, c := getName()
	fmt.Println(a, b, c)
}
