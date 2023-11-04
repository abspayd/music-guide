package main

import (
	"fmt"
)

type note struct {
	oct int 
	pitch int
}

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

func Search(pitchString string) int {
	fmt.Printf("int(pitchString[0])=%d\n", int(pitchString[0]))

	pitch := int(pitchString[0]) - int('a') // Abs(int(pitchString[0])-int('c')) ?
	fmt.Println(pitch)

	// if len(pitchString) > 1 {
	// 	if contains sharp {
	// 		pitch++
	// 	} else if contains flat {
	// 		pitch--
	// 	}
	// }
 
	return pitch
	// return BinarySearch(pitchString, asFlat, 0, len(pitchString)-1)
}

func BinarySearch(pitchString string, asFlat bool, left int, right int) int {
	// TODO: use binary search to look for pitch (look at int(pitchString[0]) as search value)

	// OR: use ascii integer values a-g and increment/decrement based on #'s or b's present
	return -1
}

func PitchIndex(pitchString string, asFlat bool) int {
	pitches := PitchArray(asFlat)
	for i := 0; i < len(pitches); i++ {
		if pitches[i] == pitchString {
			return i
		}
	}
	// Pitch not found
	return -1
}

func PitchArray(asFlat bool) []string {
	// 0  1  2  3  4  5  6  7  8  9  10 11
	// c  c# d  d# e  f  f# g  g# a  a# b
	// c  db d  eb e  f  gb g  ab a  bb b
	if asFlat {
		return []string{"c","db","d","eb","e","f","gb", "g","ab","a","bb","b"}
	}
	return []string{"c","c#","d","d#","e","f","f#", "g","g#","a","a#","b"}
}

func PitchToString(pitch int, asFlat bool) string {
	pitches := PitchArray(asFlat)

	return pitches[pitch % len(pitches)]
}

func IntervalToString(interval int) string {
	// TODO: support P8 and above
	intervals := []string{"PUNI","m2","M2","m3","M3","P4","tri","P5","m6","M6","m7","M7"}
	return intervals[interval % len(intervals)]
}
