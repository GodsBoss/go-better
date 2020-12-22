package fmt_test

import (
	"github.com/GodsBoss/go-better/fmt"

	"io"
	"os"
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
