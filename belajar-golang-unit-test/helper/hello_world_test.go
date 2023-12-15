package helper

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	resault := HelloWorld("Jhon")
	if resault != "Hello Jhon" {
		panic("Resault is not 'Hello Jhon'")
	}
}
func TestHelloWorldSmith(t *testing.T) {
	resault := HelloWorld("Smith")
	if resault != "Hello Smith" {
		panic("Resault is not 'Hello Smith'")
	}
}
