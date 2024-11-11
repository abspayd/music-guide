package app

import (
	"fmt"
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
	for i, note := range sharps {
		p, err := NewPitch(note)
		if err != nil {
			t.Error(err)
		}
		if p.index != i {
			t.Errorf("Incorrect index %d for \"%s\", expected %d", p.index, note, i)
		}
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

func TestPitchStepUpDown(t *testing.T) {
	// sharp notes (1 octave range)
	for i, note1 := range sharps {
		pitch1, err := NewPitch(note1)
		if err != nil {
			t.Error(err)
		}

		for j, note2 := range sharps {
			pitch2, err := NewPitch(note2)
			if err != nil {
				t.Error(err)
			}

			// do the interval jump
			distance := j - i

			modified_pitch := pitch1
			modified_pitch.PitchStepUpDown(distance, false)

			if modified_pitch.index != pitch2.index || modified_pitch.Class != pitch2.Class || modified_pitch.Octave != pitch2.Octave {
				t.Errorf("%v.PitchStepUpDown(%d, false) = %v, expected %v", pitch1, distance, modified_pitch, pitch2)
			}
		}
	}

	// flat notes (1 octave range)
	for i, note1 := range flats {
		pitch1, err := NewPitch(note1)
		if err != nil {
			t.Error(err)
		}

		for j, note2 := range flats {
			pitch2, err := NewPitch(note2)
			if err != nil {
				t.Error(err)
			}

			// do the interval jump
			distance := j - i

			modified_pitch := pitch1
			modified_pitch.PitchStepUpDown(distance, true)

			if modified_pitch.index != pitch2.index || modified_pitch.Class != pitch2.Class || modified_pitch.Octave != pitch2.Octave {
				t.Errorf("%v.PitchStepUpDown(%d, false) = %v, expected %v", pitch1, distance, modified_pitch, pitch2)
			}
		}
	}

	// multiple octave jump (sharps)
	for i, note1 := range sharps {
		note1 = fmt.Sprintf("%s%d", note1, 0)
		pitch1, err := NewPitch(note1)
		if err != nil {
			t.Error(err)
		}

		for j, note2 := range sharps {
			note2 = fmt.Sprintf("%s%d", note2, 1)
			pitch2, err := NewPitch(note2)
			if err != nil {
				t.Error(err)
			}

			// do the interval jump
			distance := (j - i) + 12

			modified_pitch := pitch1
			modified_pitch.PitchStepUpDown(distance, false)

			if modified_pitch.index != pitch2.index || modified_pitch.Class != pitch2.Class || modified_pitch.Octave != pitch2.Octave {
				t.Errorf("%v.PitchStepUpDown(%d, false) = %v, expected %v", pitch1, distance, modified_pitch, pitch2)
			}
		}
	}

	// multiple octave jump (flats)
	for i, note1 := range flats {
		note1 = fmt.Sprintf("%s%d", note1, 0)
		pitch1, err := NewPitch(note1)
		if err != nil {
			t.Error(err)
		}

		for j, note2 := range flats {
			note2 = fmt.Sprintf("%s%d", note2, 1)
			pitch2, err := NewPitch(note2)
			if err != nil {
				t.Error(err)
			}

			// do the interval jump
			distance := (j - i) + 12

			modified_pitch := pitch1
			modified_pitch.PitchStepUpDown(distance, true)

			if modified_pitch.index != pitch2.index || modified_pitch.Class != pitch2.Class || modified_pitch.Octave != pitch2.Octave {
				t.Errorf("%v.PitchStepUpDown(%d, false) = %v, expected %v", pitch1, distance, modified_pitch, pitch2)
			}
		}
	}
}
