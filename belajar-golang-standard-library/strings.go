package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Jhon Smith", "Jhon"))
	fmt.Println(strings.Split("Jhon Smith", " "))
	fmt.Println(strings.ToLower("Jhon Smith"))
	fmt.Println(strings.ToUpper("Jhon Smith"))
	fmt.Println(strings.Trim("Jhon Smith", " "))
	fmt.Println(strings.ReplaceAll("Jhon Smith", "Jhon", "Budi"))
}
