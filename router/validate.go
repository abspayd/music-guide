package router

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"regexp"

	"github.com/abspayd/music-companion/music"
)

var (
	// Paths found in the templates folder (./tmpl)
	validPath = regexp.MustCompile("^/(home|intervals)$")
)

// Validate the path of a request
func validatePath(path string) ([]string, error) {
	m := validPath.FindStringSubmatch(path)
	if m == nil {
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
