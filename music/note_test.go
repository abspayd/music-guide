package music

import (
	"testing"
	"math"
)

func TestSearch(t *testing.T) {
	test := map[string]int{
		// Input: expected result
		"c":   0,
		"b#":  0,
		"d##": 4,
		"e#":  5,
		"f#":  6,
		"gb":  6,
		"cbb": 10,
		"b":   11,
	}
	for note, expected := range test {
		res, _ := Search(note)
		if res != expected {
			t.Errorf("Search(\"%s\") = %d; expected %d", note, res, expected)
		}
	}

	invalid := []string{
		"",
		"h",
		"c%",
	}
	for _, note := range invalid {
		_, err := Search(note)
		if err == nil {
			t.Errorf("Search(\"%s\") did not return an error", note)
		}
	}
}

func TestGetInterval(t *testing.T) {
	octavesToTest := 3
	for j := 0; j < 12; j++ {
		n1 := Note{Pitch: j, Octave: 0} // Base to get intervals from
		for i := 0; i < 12*octavesToTest; i++ {
			n2 := Note{Pitch: i % 12, Octave: int(i / 12)}
			interval := n1.GetInterval(n2)

			if interval != int(math.Abs(float64(i-j))) {
				t.Errorf("(Note%+v).GetInterval(Note%+v) = %d; expected %d", n1, n2, interval, i)
			}
		}
	}
}
