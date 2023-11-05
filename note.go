package main

import (
	"fmt"
)	

type note struct {
	oct int 
	pitch int
}

var (
	// 	// 0  1  2  3  4  5  6  7  8  9  10 11
	// 	// c  c# d  d# e  f  f# g  g# a  a# b
	// 	// c  db d  eb e  f  gb g  ab a  bb b
	sharpNotes = []string{"c","c#","d","d#","e","f","f#", "g","g#","a","a#","b"}
	flatNotes = []string{"c","db","d","eb","e","f","gb", "g","ab","a","bb","b"}
)

func (n note) GetInterval(n2 note) int {
	interval := (n2.pitch + n2.oct * 12) - (n.pitch + n.oct * 12)

	if interval < 0 {
		interval *= -1
	}

	return interval
}

func (n note) GetNoteAtInterval(interval int) note {
	note := note{ oct: 1, pitch: 0 }
	return note
}

func (n note) PrintNote() {
	pitch := PitchToString(n.pitch, false)
	fmt.Printf("%s%d\n", pitch, n.oct)
}

func (n note) ToString() string {
	pitch := PitchToString(n.pitch, false)
	return fmt.Sprintf("%s%d", pitch, n.oct)
}

/**
* Get the index of a pitch from its string value
*/
func Search(pitchString string) int {
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
	return i
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
