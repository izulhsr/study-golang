package main

import "fmt"

func logging() {
	fmt.Println("Selesai eksekusi")
}

func runApplication() {
	defer logging()
	fmt.Println("Run Application")
}
func main() {
	runApplication()
}
