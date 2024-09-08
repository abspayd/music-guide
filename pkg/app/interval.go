package app

import ()

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
)

// Get the interval name between two pitches (given as string names)
func Interval(pitch1, pitch2 string) (string, error) {
	return "", nil
}
