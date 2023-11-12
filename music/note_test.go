package music

import (
	"testing"
)

func TestSearch(t *testing.T) {
	// Test c = 0
	noteString := "c"
	expected := 0
	i, _ := Search(noteString)
	if i != expected {
		t.Errorf("Search(\"%s\") = %d; expected %d", noteString, i, expected)
	}

	// Test b = 11
	noteString = "b"
	expected = 11
	i, _ = Search(noteString)
	if i != expected {
		t.Errorf("Search(\"%s\") = %d; expected %d", noteString, i, expected)
	}

	// Test f# = 6
	noteString = "f#"
	expected = 6
	i, _ = Search(noteString)
	if i != expected {
		t.Errorf("Search(\"%s\") = %d; expected %d", noteString, i, expected)
	}

	// Test gb = 6
	noteString = "gb"
	expected = 6
	i, _ = Search(noteString)
	if i != expected {
		t.Errorf("Search(\"%s\") = %d; expected %d", noteString, i, expected)
	}

	// Test b# = 0
	noteString = "b#"
	expected = 0
	i, _ = Search(noteString)
	if i != expected {
		t.Errorf("Search(\"%s\") = %d; expected %d", noteString, i, expected)
	}

	// Test e# = 5
	noteString = "e#"
	expected = 5
	i, _ = Search(noteString)
	if i != expected {
		t.Errorf("Search(\"%s\") = %d; expected %d", noteString, i, expected)
	}

	// Test double sharp
	noteString = "d##"
	expected = 4
	i, _ = Search(noteString)
	if i != expected {
		t.Errorf("Search(\"%s\") = %d; expected %d", noteString, i, expected)
	}

	// Test double flat
	noteString = "cbb"
	expected = 10
	i, _ = Search(noteString)
	if i != expected {
		t.Errorf("Search(\"%s\") = %d; expected %d", noteString, i, expected)
	}
}

func TestIntervalToString(t *testing.T) {
	// TODO
}
