package main

import "fmt"

func random() interface{} {
	return true
}

func main() {
	resault := random()
	switch value := resault.(type) {
	case string:
		fmt.Println(value)
	case int:
		fmt.Println(value)
	default:
		fmt.Println("Tidak Diketahui")
	}
}
