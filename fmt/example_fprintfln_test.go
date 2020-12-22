package fmt_test

import (
	"github.com/GodsBoss/go-better/fmt"

	"os"
)

func ExampleFprintfln() {
	fmt.Fprintfln(os.Stdout, "Hello, world!")
	fmt.Fprintfln(os.Stdout, "Hello, %s!", "world")
	fmt.Fprintfln(os.Stdout, "%s, %s!", "Hello", "world")

	// Output:
	// Hello, world!
	// Hello, world!
	// Hello, world!
}
