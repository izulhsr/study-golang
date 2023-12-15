package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldJhon(t *testing.T) {
	resault := HelloWorld("Jhon")
	if resault != "Hello Jhon" {
		t.Error("Resault must be Hello Jhon")
	}
	fmt.Println("Test Done")
}
func TestHelloWorldSmith(t *testing.T) {
	resault := HelloWorld("Smith")
	if resault != "Hello Smith" {
		t.Fatal("Resault must be Hello Smith")
	}
	fmt.Println("Test Done")

}
