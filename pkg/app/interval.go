package app

import (
	"fmt"
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
		13: "Minor ninth",
		14: "Major ninth",
		15: "Minor tenth",
		16: "Major tenth",
		17: "Perfect eleventh",
		18: "Augmented eleventh / Diminished twelfth",
		19: "Perfect twelfth",
		20: "Minor thirteenth",
		21: "Major thirteenth",
		22: "Minor fourteenth",
		23: "Major fourteenth",
		24: "Perfect fifteenth",
	}
)

func intervalDistance(pitch1, pitch2 Pitch) int {
	distance := (pitch2.index + (12 * pitch2.Octave)) - (pitch1.index + (12 * pitch1.Octave))
	if distance < 0 {
		distance *= -1
	}
	return distance
}

// Get the interval name between two pitch strings
func IntervalName(pitch_string_1, pitch_string_2 string) (string, error) {
	p1, err := NewPitch(pitch_string_1)
	if err != nil {
		return "", err
	}
	p2, err := NewPitch(pitch_string_2)
	if err != nil {
		return "", err
	}

	interval_distance := intervalDistance(p1, p2)
	if interval_distance >= len(interval_names) {
		return "", fmt.Errorf("Interval distance %d is greater than the maximum %d", interval_distance, len(interval_names))
	}

	return interval_names[interval_distance], nil
}
