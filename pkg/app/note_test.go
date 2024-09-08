package app

import (
	"testing"
)

func TestNewPitch(t *testing.T) {
	// Test valid pitch bases
	for c := int('A'); c <= int('G'); c++ {
		_, err := NewPitch(string(rune(c)))
		if err != nil {
			t.Error(err)
		}
	}
	// Test invalid pitch class
	_, err := NewPitch("H")
	if err == nil {
		t.Errorf("Validation did not catch invalid pitch class: \"H\"")
	}

	// Test valid accidentals
	_, err = NewPitch("c#")
	if err != nil {
		t.Error(err)
	}
	_, err = NewPitch("c##")
	if err != nil {
		t.Error(err)
	}
	_, err = NewPitch("cb")
	if err != nil {
		t.Error(err)
	}
	_, err = NewPitch("cbb")
	if err != nil {
		t.Error(err)
	}
	// Test invalid accidentals
	_, err = NewPitch("c###")
	if err == nil {
		t.Error("Validation failed to limit number of sharps")
	}
	_, err = NewPitch("cbbb")
	if err == nil {
		t.Error("Validation failed to limit number of flats")
	}

	// Test valid octaves
	_, err = NewPitch("c-2")
	if err != nil {
		t.Error(err)
	}
	_, err = NewPitch("c8")
	if err != nil {
		t.Error(err)
	}
	// Test invalid octaves
	_, err = NewPitch("c-3")
	if err == nil {
		t.Error("Validation failed to limit minimum octave")
	}
	_, err = NewPitch("c9")
	if err == nil {
		t.Error("Validation failed to limit maximum octave")
	}

	// Test indices for pitch values
	sharps := map[string]int{
		"C": 0, "C#": 1, "D": 2, "D#": 3, "E": 4, "F": 5, "F#": 6, "G": 7, "G#": 8, "A": 9, "A#": 10, "B": 11,
	}
	for k, v := range sharps {
		p, err := NewPitch(k)
		if err != nil {
			t.Error(err)
		}
		if p.index != v {
			t.Errorf("Incorrect index %d for \"%s\", expected %d", p.index, k, v)
		}
	}
	flats := map[string]int{
		"C": 0, "Db": 1, "D": 2, "Eb": 3, "E": 4, "F": 5, "Gb": 6, "G": 7, "Ab": 8, "A": 9, "Bb": 10, "B": 11,
	}
	for k, v := range flats {
		p, err := NewPitch(k)
		if err != nil {
			t.Error(err)
		}
		if p.index != v {
			t.Errorf("Incorrect index %d for \"%s\", expected %d", p.index, k, v)
		}
	}
}
