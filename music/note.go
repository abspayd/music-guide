package music

import (
	"fmt"
	"errors"
	"regexp"
)	

type Note struct {
	Octave int 
	Pitch int
}

var (
	// 	// 0  1  2  3  4  5  6  7  8  9  10 11
	// 	// c  c# d  d# e  f  f# g  g# a  a# b
	// 	// c  db d  eb e  f  gb g  ab a  bb b
	sharpNotes = []string{"c","c#","d","d#","e","f","f#", "g","g#","a","a#","b"}
	flatNotes = []string{"c","db","d","eb","e","f","gb", "g","ab","a","bb","b"}
)

func (n Note) GetInterval(n2 Note) int {
	interval := (n2.Pitch + n2.Octave * 12) - (n.Pitch + n.Octave * 12)

	if interval < 0 {
		interval *= -1
	}

	return interval
}

func (n Note) GetNoteAtInterval(interval int) Note {
	note := Note{ Octave: 1, Pitch: 0 }
	return note
}

func (n Note) PrintNote() {
	pitch := PitchToString(n.Pitch, false)
	fmt.Printf("%s%d\n", pitch, n.Octave)
}

func (n Note) ToString() string {
	pitch := PitchToString(n.Pitch, false)
	return fmt.Sprintf("%s%d", pitch, n.Octave)
}

/**
* Get the index of a pitch from its string value
*/
func Search(pitchString string) (int, error) {
	
	validPitch := regexp.MustCompile("^[a-gA-G][#,b]*$")
	m := validPitch.MatchString(pitchString)
	if !m {
		return -1, errors.New(fmt.Sprintf("Invalid pitch: %s", pitchString))
	}

	// Get values for basic notes (no sharps or flats)
	i := (int(pitchString[0] + 6) % 7) * 2
	if i > 4 {
		i--
	}
	// increment pitch by 1 for every sharp;
	// decrement pitch by 1 for every flat
	l := len(pitchString)
	if l > 1 {
		for j := 1; j < l; j++ {
			if pitchString[j] == '#' {
				i = (i+1) % 12
			} else if pitchString[j] == 'b' {
				i = (i+11) % 12 // Wrap-around subtraction
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

	return pitches[pitch % len(pitches)]
}

func IntervalToString(interval int) string {
	// TODO: support P8 and above
	intervals := []string{"PUNI","m2","M2","m3","M3","P4","tri","P5","m6","M6","m7","M7"}
	return intervals[interval % len(intervals)]
}
