package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Jhon",
			request:  "Jhon",
			expected: "Hello Jhon",
		},
		{
			name:     "Smith",
			request:  "Smith",
			expected: "Hello Smith",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Jhon", func(t *testing.T) {
		resault := HelloWorld("Jhon")
		require.Equal(t, "Hello Jhon", resault, "Resault must be Hello Jhon")
	})
	t.Run("Smith", func(t *testing.T) {
		resault := HelloWorld("Smith")
		require.Equal(t, "Hello Smith", resault, "Resault must be Hello Jhon")
	})
}

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
