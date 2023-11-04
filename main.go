package main

import (
	"fmt"
)

func main() {
	fmt.Println("== Fretboard ==")
	PrintFretBoard()

	fmt.Println()

	// Note search
	n := note{ pitch:Search("c#"), oct:0}
	n.PrintNote()
}
