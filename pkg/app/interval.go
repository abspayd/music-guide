package app

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var (
	// Intervals named by number of steps from starting pitch
	interval_names = map[int]string{
		0:  "Perfect unison",
		1:  "Minor second",
		2:  "Major second",
		3:  "Minor third",
		4:  "Major third",
		5:  "Perfect fourth",
		6:  "Tritone",
		7:  "Perfect fifth",
		8:  "Minor sixth",
		9:  "Major sixth",
		10: "Minor seventh",
		11: "Major seventh",
		12: "Perfect octave",
	}

	// Natural pitches and their index values relative to the 12 pitches
	naturals = map[rune]int{
		'C': 0,
		'D': 2,
		'E': 4,
		'F': 5,
		'G': 7,
		'A': 9,
		'B': 11,
	}

	// All note names, named as sharp
	sharp_notes = []string{
		"C",
		"C#",
		"D",
		"D#",
		"E",
		"F",
		"F#",
		"G",
		"G#",
		"A",
		"A#",
		"B",
	}

	// All note names, named as flats
	flat_notes = []string{
		"C",
		"Db",
		"D",
		"Eb",
		"E",
		"F",
		"Gb",
		"G",
		"Ab",
		"A",
		"Bb",
		"B",
	}
)

type pitch struct {
	Class  string // A note's name. E.g.: C, C#, Cb, etc.
	Octave int    // A note's octave register
}

func (p pitch) validatePitch() bool {
	if len(p.Class) == 0 {
		return false
	}

	r := regexp.MustCompile(`^[A-G][b#]{0,2}$`)
	if !r.MatchString(p.Class) {
		return false
	}
	if p.Octave < -2 || p.Octave > 8 {
		return false
	}
	return true
}

// Returns the string value for a pitch by its index
func pitchName(pitch_index int, sharp bool) (string, error) {
	if pitch_index < 0 {
		return "", fmt.Errorf("Invalid pitch index: %d. Expected a positive integer.", pitch_index)
	}

	if sharp {
		return sharp_notes[pitch_index%12], nil
	}
	return flat_notes[pitch_index%12], nil
}

func parsePitch(text string) (p pitch, err error) {
	r := regexp.MustCompile(`(?i)^(?P<class>[a-g])(?P<accidental>[b#]{0,2})(?P<octave>-?[0-9]+)?$`)

	if !r.MatchString(text) {
		return p, fmt.Errorf("Invalid pitch string: \"%s\"", text)
	}

	group_matches := r.FindStringSubmatch(text)
	group_names := r.SubexpNames()
	for idx, match := range group_matches {
		group_name := group_names[idx]
		switch group_name {
		case "class":
			p.Class += strings.ToUpper(match)
			break
		case "accidental":
			p.Class += strings.ToLower(match)
			break
		case "octave":
			if len(match) == 0 {
				break
			}
			octave, err := strconv.ParseInt(match, 10, 32)
			p.Octave = int(octave)
			if err != nil {
				return p, err
			}
			break
		default:
			break
		}
	}

	// Make sure the pitch struct we made is valid
	if !p.validatePitch() {
		return p, fmt.Errorf("Pitch is not valid: \"%v\"", p)
	}
	return p, err
}

// Gets the index value for a pitch string.
//
// E.g., c -> 0, c# -> 1, etc.
//
// Returns an error when the pitch string is not valid/known.
func pitchIndex(pitch_name string) (int, error) {
	valid_pitch_regex := regexp.MustCompile(`(?i)^[a-g][b#]{0,2}?$`)
	if !valid_pitch_regex.MatchString(pitch_name) {
		return -1, fmt.Errorf("Unexpected pitch name: \"%s\"", pitch_name)
	}

	pitch_name = strings.ToLower(pitch_name) // Make everything lower case for consistency

	base := unicode.ToUpper(rune(pitch_name[0])) // Note pitch class denoted as a single upper-case letter
	index := naturals[base]
	for _, c := range pitch_name[1:] {
		if c == 'b' {
			index = (index + 11) % 12 // Wrap-around subtraction
			continue
		}
		index = (index + 1) % 12 // Wrap-around addition
	}

	return index, nil
}

// Get the interval name between two pitches (given as string names)
func Interval(pitch1, pitch2 string) (string, error) {
	pitch1_index, err := pitchIndex(pitch1)
	if err != nil {
		return "", err
	}
	pitch2_index, err := pitchIndex(pitch2)
	if err != nil {
		return "", err
	}

	interval_distance := pitch2_index - pitch1_index

	if interval_distance < 0 {
		// Invert the interval?
		// This assumes that -12 is the minimum input value. Future to consider values above/below 12 steps.
		interval_distance += 12
	}

	name, exists := interval_names[interval_distance]

	if !exists {
		return "", fmt.Errorf("Unknown interval distance: %d", interval_distance)
	}
	return name, nil
}
