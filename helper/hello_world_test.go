package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cant run on macos")
	}

	result := HelloWorld("Zakir")
	assert.Equal(t, "Hello Zakir", result, "Result Must Be Hello Zakir")

	fmt.Println("TestHelloWorld done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Zakir")

	assert.Equal(t, "Hello Zakir", result, "Result Must Be Hello Zakir")

	fmt.Println("TestHelloWorld done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Wibu")

	require.Equal(t, "Hello Zakir", result, "Result Must Be Hello Zakir")

	fmt.Println("TestHelloWorld done")
}

func TestHelloWorldZakir(t *testing.T) {
	result := HelloWorld("Zakir")

	if result != "Hello Zakir" {
		// error cuy
		t.Fatal("Result must be Hello Zakir")
	}

	fmt.Println("TestHelloWorldZakir done")
}
