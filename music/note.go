package music

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Note struct {
	Octave int
	Pitch  int
}

var (
	// 	// 0  1  2  3  4  5  6  7  8  9  10 11
	// 	// c  c# d  d# e  f  f# g  g# a  a# b
	// 	// c  db d  eb e  f  gb g  ab a  bb b
	sharpNotes = []string{"c", "c#", "d", "d#", "e", "f", "f#", "g", "g#", "a", "a#", "b"}
	flatNotes  = []string{"c", "db", "d", "eb", "e", "f", "gb", "g", "ab", "a", "bb", "b"}
)

// Get the interval between two Notes
func (n Note) GetInterval(n2 Note) int {
	interval := (n2.Pitch + n2.Octave*12) - (n.Pitch + n.Octave*12)

	if interval < 0 {
		interval *= -1
	}

	return interval
}

// Find the note at a specific interval above the current note
func (n Note) GetNoteAtInterval(interval int) Note {
	pitch := (n.Pitch + interval) % 12
	octave := int((n.Pitch + interval) / 12) + n.Octave
	n2 := Note{
		Pitch: pitch,
		Octave: octave,
	}
	return n2
}

/**
* Get the index of a pitch from its string value
 */
func Search(pitchString string) (int, error) {
	validPitch := regexp.MustCompile("^\\s*[a-gA-G][#b]*\\s*$")
	m := validPitch.MatchString(pitchString)
	if !m {
		return -1, errors.New(fmt.Sprintf("Invalid pitch: %s", pitchString))
	}
	pitchString = strings.ToLower(strings.Trim(pitchString, " ")) // Trim whitespaces and lower case

	// Get values for basic notes (no sharps or flats)
	i := (int(pitchString[0]+6) % 7) * 2
	if i > 4 {
		i--
	}
	// increment pitch by 1 for every sharp;
	// decrement pitch by 1 for every flat
	l := len(pitchString)
	if l > 1 {
		for j := 1; j < l; j++ {
			if pitchString[j] == '#' {
				i = (i + 1) % 12
			} else if pitchString[j] == 'b' {
				i = (i + 11) % 12 // Wrap-around subtraction
			}
		}
	}
	return i, nil
}

func PitchToString(pitch int, asFlat bool) string {
	var pitches []string
	if asFlat {
		pitches = flatNotes
	} else {
		pitches = sharpNotes
	}

	return pitches[pitch%len(pitches)]
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
