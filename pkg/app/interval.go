package app

import (
	"fmt"
)

func IntervalDistance(pitch1, pitch2 int) int {
	return pitch2 - pitch1
}

func IntervalName(interval_distance int) (string, error) {

	interval_names := map[int]string{
		0: "Perfect unison",
		1: "Minor second",
		2: "Major second",
		3: "Minor third",
		4: "Major third",
		5: "Perfect fourth",
		6: "Tritone",
		7: "Perfect fifth",
		8: "Minor sixth",
		9: "Major sixth",
		10: "Minor seventh",
		11: "Major seventh",
		12: "Perfect octave",
	}

	if interval_distance < 0 {
		// Invert the interval?
		// This assumes that -12 is the minimum input value. Future to consider values above/below 12 steps.
		interval_distance += 12
	}

	name, exists := interval_names[interval_distance]

	if !exists {
		return "", fmt.Errorf("Unknown interval distance: %d", interval_distance)
	}

	fmt.Printf("%d = %s\n", interval_distance, name)
	return name, nil
}
