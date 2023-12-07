package main

import "fmt"

func factorial(number int) int {
	hasil := 1
	for i := number; i > 0; i-- {
		hasil *= i
	}

	return hasil
}

func factorialNumber(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialNumber(value-1)
	}
}

func main() {
	fmt.Println(factorial(10))
	fmt.Println(factorialNumber(10))
}
