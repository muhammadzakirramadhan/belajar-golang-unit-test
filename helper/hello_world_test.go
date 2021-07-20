package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchMarkSub(b *testing.B) {
	b.Run("Zakir", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Zakir")
		}
	})
}

func BenchMarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Zakir")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Zakir",
			request:  "Zakir",
			expected: "Hello Zakir",
		}, {
			name:     "Wibu",
			request:  "Wibu",
			expected: "Hello Wibu",
		}, {
			name:     "Azzah",
			request:  "Azzah",
			expected: "Hello Azzah",
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
	t.Run("Zakir", func(t *testing.T) {
		result := HelloWorld("Zakir")
		require.Equal(t, "Hello Zakir", result, "Result Must Be Hello Zakir")
	})

	t.Run("Wibu", func(t *testing.T) {
		result := HelloWorld("Wibu")
		require.Equal(t, "Hello Wibu", result, "Result Must Be Hello Zakir")
	})
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cant run on macos")
	}

	result := HelloWorld("Zakir")
	assert.Equal(t, "Hello Zakir", result, "Result Must Be Hello Zakir")

	fmt.Println("TestHelloWorld done")
}

func TestMain(m *testing.M) {

	// before
	fmt.Println("Sebelum Run")

	m.Run()

	// after
	fmt.Println("Setelah run")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Zakir")

	assert.Equal(t, "Hello Zakir", result, "Result Must Be Hello Zakir")

	fmt.Println("TestHelloWorld done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Zakir")

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
