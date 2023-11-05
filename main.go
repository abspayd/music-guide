package main

import (
	"fmt"
)

func main() {
	// fmt.Println("== Fretboard ==")
	// PrintFretBoard()

	// fmt.Println()

	p:="c"
	fmt.Printf("%s=%d\n", p, Search(p))
	p="d"
	fmt.Printf("%s=%d\n", p, Search(p))
	p="e"
	fmt.Printf("%s=%d\n", p, Search(p))
	p="f"
	fmt.Printf("%s=%d\n", p, Search(p))
	p="g"
	fmt.Printf("%s=%d\n", p, Search(p))
	p="a"
	fmt.Printf("%s=%d\n", p, Search(p))
	p="b"
	fmt.Printf("%s=%d\n", p, Search(p))
}
