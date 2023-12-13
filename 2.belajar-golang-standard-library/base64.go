package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	value := "Jhon Smith Lance"
	encode := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encode)
	decode, err := base64.StdEncoding.DecodeString(encode)
	if err != nil {
		fmt.Println("error", err.Error())
	} else {
		fmt.Println(string(decode))
	}
}
