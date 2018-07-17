package main

import humanize "github.com/dustin/go-humanize"
import "fmt"

func main() {
	fmt.Printf("Hola Mundo\n")
	fmt.Printf("That file is %s.\n", humanize.Bytes(82854982))
	// testNonblock()
}
