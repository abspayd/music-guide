package app

import (
	"testing"
)

func TestInterval(t *testing.T) {
	pitch1 := "c"
	pitch2 := "c#"
	expected := interval_names[1]
	result, err := Interval(pitch1, pitch2)
	if err != nil {
		t.Error(err)
	}
	if expected != result {
		t.Errorf("Interval(\"%s\", \"%s\") = \"%s\", expected: \"%s\"\n", pitch1, pitch2, result, expected)
	}

	pitch1 = "e"
	pitch2 = "b"
	expected = interval_names[7]
	result, err = Interval(pitch1, pitch2)
	if err != nil {
		t.Error(err)
	}
	if expected != result {
		t.Errorf("Interval(\"%s\", \"%s\") = \"%s\", expected: \"%s\"\n", pitch1, pitch2, result, expected)
	}
	
	_, err = parsePitch("c#-1")
	if err != nil {
		t.Error(err)
	}
	_, err = parsePitch("c##-2")
	if err != nil {
		t.Error(err)
	}
	_, err = parsePitch("c10")
	if err != nil {
		t.Error(err)
	}
	_, err = parsePitch("a")
	if err != nil {
		t.Error(err)
	}
	_, err = parsePitch("-")
	if err != nil {
		t.Error(err)
	}
	_, err = parsePitch("a-")
	if err != nil {
		t.Error(err)
	}
	t.Fail()
}

func TestPitchIndex(t *testing.T) {
	pitch := "c"
	index, err := pitchIndex(pitch)
	expected_index := 0
	if err != nil {
		t.Error(err)
	}
	if index != expected_index {
		t.Errorf("pitchIndex(\"%s\") = %d, expected: %d\n", pitch, index, expected_index)
	}
	pitch = "C"
	index, err = pitchIndex(pitch)
	expected_index = 0
	if err != nil {
		t.Error(err)
	}
	if index != expected_index {
		t.Errorf("pitchIndex(\"%s\") = %d, expected: %d\n", pitch, index, expected_index)
	}

	pitch = "cbb"
	index, err = pitchIndex(pitch)
	expected_index = 10
	if err != nil {
		t.Error(err)
	}
	if index != expected_index {
		t.Errorf("pitchIndex(\"%s\") = %d, expected: %d\n", pitch, index, expected_index)
	}
	pitch = "c##"
	index, err = pitchIndex(pitch)
	expected_index = 2
	if err != nil {
		t.Error(err)
	}
	if index != expected_index {
		t.Errorf("pitchIndex(\"%s\") = %d, expected: %d\n", pitch, index, expected_index)
	}
	pitch = "b##"
	index, err = pitchIndex(pitch)
	expected_index = 1
	if err != nil {
		t.Error(err)
	}
	if index != expected_index {
		t.Errorf("pitchIndex(\"%s\") = %d, expected: %d\n", pitch, index, expected_index)
	}
}

func TestPitchValidation(t *testing.T) {
	// Invalid inputs fail validation
	pitch := "h"
	_, err := pitchIndex(pitch)
	if err == nil {
		t.Errorf("Regex failed to validate invalid pitch: \"%s\"", pitch)
	}
	pitch = "cbbb"
	_, err = pitchIndex(pitch)
	if err == nil {
		t.Error("Validation failed to catch too many flats (b's)")
	}
	pitch = "c###"
	_, err = pitchIndex(pitch)
	if err == nil {
		t.Error("Validation failed to catch too many sharps (#'s)")
	}

	// Valid inputs pass validation
	for i := int('a'); i <= int('g'); i++ {
		// Naturals
		c := string(rune(i))
		_, err = pitchIndex(c)
		if err != nil {
			t.Errorf("Validation failed valid pitch: \"%s\"", pitch)
		}
		// Sharps
		sharp := c + "#"
		_, err = pitchIndex(sharp)
		if err != nil {
			t.Errorf("Validation failed valid pitch: \"%s\"", sharp)
		}
		// Double sharps
		sharp += "#"
		_, err = pitchIndex(sharp)
		if err != nil {
			t.Errorf("Validation failed valid pitch: \"%s\"", sharp)
		}

		// Flats
		flat := c + "b"
		_, err = pitchIndex(flat)
		if err != nil {
			t.Errorf("Validation failed valid pitch: \"%s\"", flat)
		}
		// Double flats
		flat += "b"
		_, err = pitchIndex(flat)
		if err != nil {
			t.Errorf("Validation failed valid pitch: \"%s\"", flat)
		}
	}
}
