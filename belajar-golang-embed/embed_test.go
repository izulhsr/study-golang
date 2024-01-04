package belajargolangembed

import (
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var embed string

func TestString(t *testing.T) {
	fmt.Println(embed)
}

//go:embed logo.png
var logo []byte

func TestLogo(t *testing.T) {
	err := ioutil.WriteFile("logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
