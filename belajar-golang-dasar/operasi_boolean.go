package main

import "fmt"

func main() {
	nilaiUjian := 90
	nilaiAbsen := 80

	lulusNilaiUjian := nilaiUjian > 80
	lulusNilaiAbsen := nilaiAbsen > 75
	lulusUjian := lulusNilaiUjian && lulusNilaiAbsen
	fmt.Println(lulusUjian)
}
