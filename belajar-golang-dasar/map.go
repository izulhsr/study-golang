package main

import "fmt"

func main() {
	person := map[string]string{
		"Nama":   "Jhon",
		"Alamat": "Boston",
	}
	fmt.Println(person)
	fmt.Println(person["Nama"])
	fmt.Println(person["Alamat"])

	book := make(map[string]string)
	book["Title"] = "Buku Golang"
	book["Author"] = "Jhon"
	book["Salah"] = "Ups"

	fmt.Println(book)
	delete(book, "Salah")
	fmt.Println(book)

}
