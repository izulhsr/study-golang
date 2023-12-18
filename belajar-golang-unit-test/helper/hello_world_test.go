package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Before Test")
	m.Run()
	fmt.Println("After Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Tdk jalan di windows")
	}
}

func TestHelloWorldRequire(t *testing.T) {
	resault := HelloWorld("Jhon")
	require.Equal(t, "Hello Jhon", resault, "Resault must be Hello Jhon")
	fmt.Println("Test Require Done")
}

func TestHelloWorldAssert(t *testing.T) {
	resault := HelloWorld("Jhon")
	assert.Equal(t, "Hello Jhon", resault, "Resault must be Hello Jhon")
	fmt.Println("Test Assert Done")
}

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
