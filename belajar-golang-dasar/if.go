package main

import "fmt"

func main() {
	name := "Budi"
	if name == "Jhon" {
		fmt.Println("Hello Jhon")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("Boleh kenalan?")
	}

	nama := "Budi"
	if panjang := len(nama); panjang > 5 {
		fmt.Println("Nama lebih dari 5 karakter")
	} else {
		fmt.Println("Nama pendek")
	}
}
