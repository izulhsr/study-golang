package main

import "fmt"

func main() {
	a := 5
	b := 10
	c := a < b
	fmt.Println("A < B", c)
	d := a > b
	fmt.Println("A > B", d)
	e := a <= b
	fmt.Println("A <= b", e)
	f := a >= b
	fmt.Println("A >= B", f)
	g := a == b
	fmt.Println("A == B", g)
	nama := "Jhon"
	nama1 := "Smith"
	fmt.Println("Nama", nama != nama1)
}
