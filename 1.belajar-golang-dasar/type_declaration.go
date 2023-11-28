package main

import "fmt"

func main() {
	type noKTP string

	var ktp noKTP = "111111111"
	fmt.Println(ktp)
	fmt.Println(noKTP("222222222"))
}
