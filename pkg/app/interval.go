package app

import (
	"fmt"
)

var (
	// Intervals named by number of steps from starting pitch
	interval_names = map[string]int{
		"Perfect unison": 0,
		"Minor second": 1,
		"Major second": 2,
		"Minor third": 3,
		"Major third": 4,
		"Perfect fourth": 5,
		"Tritone": 6,
		"Perfect fifth": 7,
		"Minor sixth": 8,
		"Major sixth": 9,
		"Minor seventh": 10,
		"Major seventh": 11,
		"Perfect octave": 12,
		"Minor ninth": 13,
		"Major ninth": 14,
		"Minor tenth": 15,
		"Major tenth": 16,
		"Perfect eleventh": 17,
		"Augmented eleventh / Diminished twelfth": 18,
		"Perfect twelfth": 19,
		"Minor thirteenth": 20,
		"Major thirteenth": 21,
		"Minor fourteenth": 22,
		"Major fourteenth": 23,
		"Perfect fifteenth": 24,
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

	for k, v := range interval_names {
		if v == interval_distance {
			return k, nil
		}
	}

	return "", fmt.Errorf("An error occurred finding the interval name for interval with distance %d", interval_distance)
}

func IntervalNames() []string {
	intervals := make([]string, len(interval_names))
	for key, _ := range interval_names {
		intervals = append(intervals, key)
	}
	return intervals
}

func intervalIndex(interval_name string) (int, error) {
	name, ok := interval_names[interval_name]
	if !ok {
		return -1, fmt.Errorf("Unable to get index for interval \"%s\"", interval_name)
	}
	return name, nil
}

func NoteFromInterval(pitch Pitch, interval string) (Pitch, error) {
	interval_distance, err := intervalIndex(interval)
	if err != nil {
		return Pitch{}, err
	}
	pitch2 := pitch
	pitch2.PitchStepUpDown(interval_distance, false)
	return pitch2, nil
}
