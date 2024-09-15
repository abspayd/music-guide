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
	sharps := []string{
		"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B",
	}
	for i, note := range sharps {
		p, err := NewPitch(note)
		if err != nil {
			t.Error(err)
		}
		if p.index != i {
			t.Errorf("Incorrect index %d for \"%s\", expected %d", p.index, note, i)
		}
	}
	flats := []string{
		"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B",
	}
	for i, note := range flats {
		p, err := NewPitch(note)
		if err != nil {
			t.Error(err)
		}
		if p.index != i {
			t.Errorf("Incorrect index %d for \"%s\", expected %d", p.index, note, i)
		}
	}
}
