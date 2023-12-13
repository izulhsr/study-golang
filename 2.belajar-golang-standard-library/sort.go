package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

// Kontrak Untuk user slice
func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[i].Age
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User{
		{"Jhon", 34},
		{"Jhona", 32},
		{"Jhonc", 34},
		{"Jhonf", 24},
		{"Jhons", 44},
	}
	sort.Sort(UserSlice(users))

	fmt.Print(users)
}
