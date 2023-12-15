package main

import "fmt"

func getGuudBye(name string) string {
	return "Guud Bye " + name
}
func main() {
	firstName := getGuudBye
	fmt.Println(firstName("Jhon"))
}
