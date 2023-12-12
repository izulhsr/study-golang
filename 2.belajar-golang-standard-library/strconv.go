package main

import (
	"fmt"
	"strconv"
)

func main() {
	resultStr, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println(err, err.Error())
	} else {
		fmt.Println(resultStr)
	}

	resultInt, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println(err, err.Error())
	} else {
		fmt.Println(resultInt)
	}

	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary)

	var stringInt string = strconv.Itoa(999)
	fmt.Println(stringInt)
}
