package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	// recover benar
	message := recover()
	fmt.Println("Pesan Erorr", message)
}
func runApp(error bool) {
	defer endApp()
	if error {
		panic("Yah Error")
	}
	// Recover salah
	// message := recover()
	// fmt.Println("Pesan Erorr", message)
}

func main() {
	runApp(true)
}
