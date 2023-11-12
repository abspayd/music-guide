package music

import (
	"testing"
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

	_, err := Search("h")
	if err == nil {
		t.Errorf("Search(\"h\") did not return an error")
	}
}

func TestGetInterval(t *testing.T) {
	c := Note{Pitch: 0, Octave: 0}
	test := map[Note]int{}

	for note, expected := range test {
		c.GetInterval(note)
		_ = expected
	}
}
