package app

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	naturals = map[string]int{
		"C": 0,
		"D": 2,
		"E": 4,
		"F": 5,
		"G": 7,
		"A": 9,
		"B": 11,
	}
)

type Pitch struct {
	Class  string
	Octave int
	index  int
}

// Create a new pitch from a string
func NewPitch(str string) (Pitch, error) {
	r := regexp.MustCompile(`(?i)^(?P<class_base>[a-g])(?P<class_accidental>[b#]{0,2})(?P<octave>-?\d)?$`)

	p := Pitch{}

	if !r.MatchString(str) {
		return Pitch{}, fmt.Errorf("Invalid pitch: \"%s\"", str)
	}

	// Get the pitch class and octave
	groups := r.SubexpNames()
	matches := r.FindStringSubmatch(str)
	p.index = 0
	for i, m := range matches {
		if i == 0 {
			// First match is the whole string
			continue
		}

		group := groups[i]
		switch group {
		case "class_base":
			p.Class = strings.ToUpper(m)
			p.index = naturals[p.Class]
			break
		case "class_accidental":
			if len(m) == 0 {
				break
			}
			accidental := strings.ToLower(m)
			p.Class += accidental
			for _, c := range accidental {
				symbol := rune(c)
				if symbol == rune('#') {
					p.index = (p.index + 1) % 12
				} else if symbol == rune('b') {
					p.index = (p.index + 11) % 12
				} else {
					return Pitch{}, fmt.Errorf("Unknown accidental symbol: '%c'", symbol)
				}
			}

			if p.index < -2 || p.index > 11 {
				return Pitch{}, fmt.Errorf("Invalid pitch index: %d", p.index)
			}
			break
		case "octave":
			if len(m) == 0 {
				p.Octave = 0
				break
			}
			octave, err := strconv.ParseInt(m, 10, 32)
			if err != nil {
				return Pitch{}, err
			}
			p.Octave = int(octave)
			if p.Octave < -2 || p.Octave > 8 {
				return Pitch{}, fmt.Errorf("Invalid octave: %d", p.Octave)
			}
			break
		default:
			break
		}
	}
	return p, nil
}

// Convert a pitch to a string
func (p Pitch) ToString() string {
	return fmt.Sprintf("%s%d", p.Class, p.Octave)
}
