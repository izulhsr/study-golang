package main

import (
	"fmt"
	"testing"
	"time"
)

func HelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutines(t *testing.T) {
	go HelloWorld()
	fmt.Println("ups")
	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("display", number)
}

func TestDisplayNumber(t *testing.T) {
	for i := 0; i < 100000; i++ {
		DisplayNumber(i)
	}
	time.Sleep(5 * time.Second)
}
