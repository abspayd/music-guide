package music

import (
	"fmt"
)

const NUMBER_OF_STRINGS = 6
const NUMBER_OF_FRETS = 20

// TODO: support alternate tunings
func Fretboard() [NUMBER_OF_STRINGS][NUMBER_OF_FRETS]Pitch {
	strings := [NUMBER_OF_STRINGS][NUMBER_OF_FRETS]Pitch{}

	// Instantiate each string with notes
	for i := 0; i < NUMBER_OF_FRETS; i++ {
		strings[0][i] = Pitch{Octave: 2, Note: i + 4%12}	// E ...
		strings[1][i] = Pitch{Octave: 1, Note: i + 11%12}	// B ...
		strings[2][i] = Pitch{Octave: 1, Note: i + 7%12}	// G ...
		strings[3][i] = Pitch{Octave: 0, Note: i + 2%12}	// D ...
		strings[4][i] = Pitch{Octave: 0, Note: i + 9%12}	// A ...
		strings[5][i] = Pitch{Octave: 0, Note: i + 4%12}	// E ...
	}

	return strings
}

func PrintFretBoard() {
	strings := Fretboard()

	for j := 0; j < 20; j++ {
		fmt.Print(j, "\t")
	}
	fmt.Println()

	for i := 0; i < 6; i++ {
		for j := 0; j < 20; j++ {
			fmt.Print(NoteToString(strings[i][j].Note, false), "\t")
		}
		fmt.Println()
	}
	fmt.Println()
}

func Locate(n Pitch) [NUMBER_OF_STRINGS]int {
	// fretboard := Fretboard()
	for i := 0; i < NUMBER_OF_STRINGS; i++ {
		// string := fretboard[i]
		// Search for note on fretboard
	}
	return [NUMBER_OF_STRINGS]int{}
}
