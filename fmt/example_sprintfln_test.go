package fmt_test

import (
	"io"
	"os"

	"github.com/GodsBoss/go-better/fmt"
)

func ExampleSprintfln() {
	io.WriteString(os.Stdout, fmt.Sprintfln("Hello, world!"))
	io.WriteString(os.Stdout, fmt.Sprintfln("Hello, %s!", "world"))
	io.WriteString(os.Stdout, fmt.Sprintfln("%s, %s!", "Hello", "world"))

	// Output:
	// Hello, world!
	// Hello, world!
	// Hello, world!
}
