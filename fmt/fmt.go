// Package fmt contains helpers representing the pinnacle of usefulness.
package fmt

import (
	"fmt"
	"io"
)

// Fprintfln is like fmt.Fprintf, but adds a newline to format.
func Fprintfln(w io.Writer, format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(w, format+"\n", a...)
}

// Printfln is like fmt.Printf, but adds a newline to format.
func Printfln(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format+"\n", a...)
}

// Sprintfln is like fmt.Sprintf, but adds a newline to format.
func Sprintfln(format string, a ...interface{}) string {
	return fmt.Sprintf(format+"\n", a...)
}
