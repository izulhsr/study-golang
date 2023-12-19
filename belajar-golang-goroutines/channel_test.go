package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	chanel := make(chan string)
	defer close(chanel)
	go func() {
		time.Sleep(2 * time.Second)
		chanel <- "Jhon Smith Lance"
		fmt.Println("Selesai Mengirim Data ke Chanel")
	}()
	data := <-chanel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func GiveMeRespond(chanel chan string) {
	time.Sleep(2 * time.Second)
	chanel <- "Give me Respond Jhon Smith Lance"
}

func TestChannelAsParameter(t *testing.T) {
	chanel := make(chan string)
	defer close(chanel)
	go GiveMeRespond(chanel)
	data := <-chanel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func OnlyIn(chanel chan<- string) {
	time.Sleep(2 * time.Second)
	chanel <- "Only In"
}

func OnlyOut(chanel <-chan string) {
	data := <-chanel
	fmt.Println(data)
}

func TestInOutChanel(t *testing.T) {
	chanel := make(chan string)
	defer close(chanel)
	go OnlyIn(chanel)
	go OnlyOut(chanel)
	fmt.Println(chanel)
	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	chanel := make(chan string, 3)
	defer close(chanel)

	go func() {
		chanel <- "Jhon"
		chanel <- "Smith"
		chanel <- "Lance"
	}()
	go func() {
		fmt.Println(<-chanel)
		fmt.Println(<-chanel)
		fmt.Println(<-chanel)
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}
