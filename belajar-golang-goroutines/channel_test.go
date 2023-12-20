package main

import (
	"fmt"
	"strconv"
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

func TestRangeChannel(t *testing.T) {
	chanel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			chanel <- "Perulangan Ke" + strconv.Itoa(i)
		}
		close(chanel)
	}()

	for data := range chanel {
		fmt.Println("Menerima Data", data)
	}
}

func TestSelectChannel(t *testing.T) {
	chanel1 := make(chan string)
	chanel2 := make(chan string)
	defer close(chanel1)
	defer close(chanel2)
	go GiveMeRespond(chanel1)
	go GiveMeRespond(chanel2)

	counter := 0
	for {
		select {
		case data := <-chanel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-chanel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelectChannel(t *testing.T) {
	chanel1 := make(chan string)
	chanel2 := make(chan string)
	defer close(chanel1)
	defer close(chanel2)
	go GiveMeRespond(chanel1)
	go GiveMeRespond(chanel2)

	counter := 0
	for {
		select {
		case data := <-chanel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-chanel2:
			fmt.Println("Data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}
		if counter == 2 {
			break
		}
	}
}