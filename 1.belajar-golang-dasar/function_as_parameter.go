package main

import "fmt"

type Filer func(string) string

func sayHelloWithFilter(name string, filter Filer) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "ajg" {
		return "***"
	} else {
		return name
	}
}

func main() {

	sayHelloWithFilter("ajg", spamFilter)
}
