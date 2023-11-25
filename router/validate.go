package router

import (
	"errors"
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/abspayd/music-companion/music"
)

var (
	// Map of valid URL paths and their valid subpaths
	pathMap = map[string][]string{
		"intervals": {
			"validation",
			"octaveModeToggle",
		},
		"home": {
			// No subpaths
		},
	}
)

// Validate the path of a request
func validatePath(path string) ([]string, error) {
	m := strings.Split(path, "/")
	if m == nil {
		return nil, errors.New("Invalid path")
	}

	if len(m) > 1 {
		// remove the first element of m, which is an empty string
		m = m[1:]
	}

	// Validate parent path
	if _, ok := pathMap[m[0]]; !ok {
		return nil, errors.New("Invalid path")
	}

	// Validate subpath
	if len(m) > 1 && slices.Contains(pathMap[m[0]], path) {
		return nil, errors.New("Invalid path")
	}

	return m, nil
}

// Handle note validation requests for the intervals page
func handleValidateNote(w http.ResponseWriter, r *http.Request, id int) {
	// Validate the note
	inputName := fmt.Sprintf("note%d", id)
	note := strings.Trim(r.FormValue(inputName), " ")
	_, err := music.Search(note)
	input := &InputField{
		Id:    id,
		Value: note,
		Error: "",
	}

	if err != nil || len(note) > 4 {
		input.Error = "Invalid note"
		templates.ExecuteTemplate(w, "invalidBasicNoteInput", input)
	} else {
		// Passed validation, restore the input
		templates.ExecuteTemplate(w, "basicNoteInput", input)
	}
}

// Handle octave validation requests for the intervals page
func handleValidateOctave(w http.ResponseWriter, r *http.Request, id int) {
	inputName := fmt.Sprintf("octave%d", id)
	inputString := r.FormValue(inputName)

	input := &InputField{
		Id:    id,
		Value: inputString,
		Error: "",
	}
	octave, err := strconv.Atoi(inputString)
	if err != nil || octave < 0 || octave > 2 {
		input.Error = "Invalid octave"
		templates.ExecuteTemplate(w, "invalidOctave", input)
	} else {
		// Passed validation, restore the input
		templates.ExecuteTemplate(w, "octaveInput", input)
	}
}
