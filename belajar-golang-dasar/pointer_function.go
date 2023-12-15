package main

import "fmt"

type Address struct{ City, Province, Country string }

func ChanceCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{}
	ChanceCountryToIndonesia(&address)

	fmt.Println(address)
}
