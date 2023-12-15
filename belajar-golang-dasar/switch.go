package main

import "fmt"

func main() {
	name := "Jhon"
	switch name {
	case "Jhon":
		fmt.Println("Hi Jhon")
	case "Budi":
		fmt.Println("Hi Budi")
	default:
		fmt.Println("Boleh Kenalan?")
	}

	switch panjang := len(name); panjang > 5 {
	case true:
		fmt.Println(len(name), "karakter")
	case false:
		fmt.Println(len(name), "karakter")
	}
}
