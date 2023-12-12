package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Jhon")
	data.PushBack("Smith")
	data.PushBack("Lance")

	for i := data.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
