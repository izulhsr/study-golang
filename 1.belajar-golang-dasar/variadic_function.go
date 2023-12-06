package main

import "fmt"

func sumAll(number ...int) int {
	total := 0
	for _, number := range number {
		total += number

	}
	return total
}
func main() {
	fmt.Println(sumAll(10, 10, 10, 10))

	angka := []int{10, 10, 10, 10, 10, 10}
	fmt.Println(sumAll(angka...))
}
