package fmt_test

import (
	"github.com/GodsBoss/go-better/fmt"
)

func ExamplePrintfln() {
	fmt.Printfln("Hello, world!")
	fmt.Printfln("Hello, %s!", "world")
	fmt.Printfln("%s, %s!", "Hello", "world")

	// Output:
	// Hello, world!
	// Hello, world!
	// Hello, world!
}
