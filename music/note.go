package music

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Pitch struct {
	Octave int
	Note  int
}

var (
	// 	// 0  1  2  3  4  5  6  7  8  9  10 11
	// 	// c  c# d  d# e  f  f# g  g# a  a# b
	// 	// c  db d  eb e  f  gb g  ab a  bb b
	sharpNotes = []string{"c", "c#", "d", "d#", "e", "f", "f#", "g", "g#", "a", "a#", "b"}
	flatNotes  = []string{"c", "db", "d", "eb", "e", "f", "gb", "g", "ab", "a", "bb", "b"}
)

// Get the interval between two Notes
func (p Pitch) GetInterval(p2 Pitch) int {
	interval := (p2.Note + p2.Octave*12) - (p.Note + p.Octave*12)

	if interval < 0 {
		interval *= -1
	}

	return interval
}

// Find the pitch at a specific interval above the current note
func (p Pitch) GetPitchAtInterval(interval int) Pitch {
	note := (p.Note + interval) % 12
	octave := int((p.Note + interval) / 12) + p.Octave
	p2 := Pitch{
		Note: note,
		Octave: octave,
	}
	return p2
}

/**
* Get the index of a pitch from its string value
 */
func Search(noteString string) (int, error) {
	validPitch := regexp.MustCompile("^\\s*[a-gA-G][#b]*\\s*$")
	m := validPitch.MatchString(noteString)
	if !m {
		return -1, errors.New(fmt.Sprintf("Invalid pitch: %s", noteString))
	}
	noteString = strings.ToLower(strings.Trim(noteString, " ")) // Trim whitespaces and lower case

	// Get values for basic notes (no sharps or flats)
	i := (int(noteString[0]+6) % 7) * 2
	if i > 4 {
		i--
	}
	// increment pitch by 1 for every sharp;
	// decrement pitch by 1 for every flat
	l := len(noteString)
	if l > 1 {
		for j := 1; j < l; j++ {
			if noteString[j] == '#' {
				i = (i + 1) % 12
			} else if noteString[j] == 'b' {
				i = (i + 11) % 12 // Wrap-around subtraction
			}
		}
	}
	return i, nil
}

func NoteToString(note int, asFlat bool) string {
	var notes []string
	if asFlat {
		notes = flatNotes
	} else {
		notes = sharpNotes
	}

	return notes[note%len(notes)]
}

func IntervalToString(interval int) string {
	intervals := []string{
		"Perfect Unison",
		"Minor second",
		"Major second",
		"Minor third",
		"Major third",
		"Perfect fourth",
		"Tritone",
		"Perfect fifth",
		"Minor sixth",
		"Major sixth",
		"Minor seventh",
		"Major seventh",
		"Perfect Octave",
		"Minor ninth",
		"Major ninth",
		"Minor tenth",
		"Major tenth",
		"Perfect eleventh",
		"Tritone",
		"Perfect twelfth",
		"Minor thirteenth",
		"Major thirteenth",
		"Minor fourteenth",
		"Major fourteenth",
		"15va",
	}

	return intervals[interval%len(intervals)]
}
