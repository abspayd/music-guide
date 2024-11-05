package app

import (
	"fmt"
)

var (
	// Intervals named by number of steps from starting pitch
	interval_names = []string{
		"Perfect unison",
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
		"Perfect octave",
		"Minor ninth",
		"Major ninth",
		"Minor tenth",
		"Major tenth",
		"Perfect eleventh",
		"Augmented eleventh / Diminished twelfth",
		"Perfect twelfth",
		"Minor thirteenth",
		"Major thirteenth",
		"Minor fourteenth",
		"Major fourteenth",
		"Perfect fifteenth",
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
func IntervalName(pitch1, pitch2 Pitch) (string, error) {
	interval_distance := intervalDistance(pitch1, pitch2)
	if interval_distance >= len(interval_names) {
		return "", fmt.Errorf("Interval distance %d is greater than the maximum %d", interval_distance, len(interval_names))
	}

	return interval_names[interval_distance], nil
}
