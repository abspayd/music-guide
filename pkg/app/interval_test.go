package app

import (
	"fmt"
	"testing"
)

func TestIntervalDistance(t *testing.T) {
	notes := []string{
		"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B",
	}

	// Test that the interval distance from C0 to B0 is just 0-11
	pitch1, err := NewPitch(notes[0])
	if err != nil {
		t.Error(err)
	}
	for i, note := range notes {
		pitch2, err := NewPitch(note)
		if err != nil {
			t.Error(err)
		}
		interval_distance := intervalDistance(pitch1, pitch2)
		if interval_distance != i {
			t.Errorf("intervalDistance(%v, %v) = %d, expected %d", pitch1, pitch2, interval_distance, i)
		}
	}

	// Test that the interval distance from C0 to C1 through B1 is 12-23
	for i, note := range notes {
		pitch2, err := NewPitch(fmt.Sprintf("%s%d", note, 1))
		if err != nil {
			t.Error(err)
		}
		interval_distance := intervalDistance(pitch1, pitch2)
		if interval_distance != (i + 12) {
			t.Errorf("intervalDistance(%v, %v) = %d, expected %d", pitch1, pitch2, interval_distance, i)
		}
	}

	// Test intervals going backwards
	pitch1, err = NewPitch("F0")
	if err != nil {
		t.Error(err)
	}
	pitch2, err := NewPitch("C#0")
	if err != nil {
		t.Error(err)
	}
	interval_distance := intervalDistance(pitch1, pitch2)
	expected_distance := 4
	if interval_distance != expected_distance {
		t.Errorf("intervalDistance(%v, %v) = %d, expected %d", pitch1, pitch2, interval_distance, expected_distance)
	}
}

func TestIntervalName(t * testing.T) {
	pitch1, err := NewPitch("E0")
	if err != nil {
		t.Error(err)
	}
	pitch2, err := NewPitch("A0")
	if err != nil {
		t.Error(err)
	}
	interval_name, err := IntervalName(pitch1, pitch2)
	if err != nil {
		t.Error(err)
	}
	expected := interval_names[5]
	if interval_name != expected {
		t.Errorf("IntervalName(\"%v\", \"%v\") = \"%s\", expected \"%s\"", pitch1, pitch2, interval_name, expected)
	}

	// Same octave but reversed is the same interval
	interval_name, err = IntervalName(pitch2, pitch1)
	if err != nil {
		t.Error(err)
	}
	expected = interval_names[5]
	if interval_name != expected {
		t.Errorf("IntervalName(\"%v\", \"%v\") = \"%s\", expected \"%s\"", pitch1, pitch2, interval_name, expected)
	}

	pitch1, err = NewPitch("A0")
	if err != nil {
		t.Error(err)
	}
	pitch2, err = NewPitch("E1")
	if err != nil {
		t.Error(err)
	}
	interval_name, err = IntervalName(pitch1, pitch2)
	expected = interval_names[7]
	if err != nil {
		t.Error(err)
	}
	if interval_name != expected {
		t.Errorf("IntervalName(\"%v\", \"%v\") = \"%s\", expected \"%s\"", pitch1, pitch2, interval_name, expected)
	}

	pitch1, err = NewPitch("A0")
	if err != nil {
		t.Error(err)
	}
	pitch2, err = NewPitch("D2")
	if err != nil {
		t.Error(err)
	}
	interval_name, err = IntervalName(pitch1, pitch2)
	expected = interval_names[17]
	if err != nil {
		t.Error(err)
	}
	if interval_name != expected {
		t.Errorf("IntervalName(\"%v\", \"%v\") = \"%s\", expected \"%s\"", pitch1, pitch2, interval_name, expected)
	}
}
