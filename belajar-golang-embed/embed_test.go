package belajargolangembed

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed version.txt
var embed string

func TestString(t *testing.T) {
	fmt.Println(embed)
}
