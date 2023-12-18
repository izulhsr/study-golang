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
