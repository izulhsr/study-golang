package main

import "fmt"

func main() {

	for counter := 0; counter <= 10; counter++ {
		fmt.Println("perulangan ke", counter)
	}
	names := []string{"Jhon", "Smith", "Lance"}
	for index, names := range names {
		fmt.Println(index, names)
	}
}
