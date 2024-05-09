package main

import (
	"fmt"
	"io"
	"os"
)

/*
// The problem: How to teet this?
func Greet(name string) {
	// Printf usa Fprintf, para el cual el primer arguemnto es un io.Writer
	// Por lo tanto el primer parámetro satisface la interfaz Writer
	fmt.Printf("Hello %s", name)
}

func main() {
	Greet("Pepe")
}
*/

// Solution 1
func Greet(writer io.Writer, name string) {
	_, _ = fmt.Fprintf(writer, "Hello %s", name)
}

func main() {
	Greet(os.Stdout, "Pepe")
}
