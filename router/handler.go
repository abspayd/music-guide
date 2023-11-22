package router

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"regexp"
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

	validPath = validPathMatches(pathMap)
	templates = loadTemplates()
)

type InputField struct {
	Id    int
	Value string
	Error string
}

// Build regexp to match valid paths
func validPathMatches(paths map[string][]string) *regexp.Regexp {
	r := `^/(`

	var roots string
	for root := range pathMap {
		roots += root + "|"
	}
	roots = roots[0 : len(roots)-1] // trim trailing "|"
	r += roots + `)(/\w+)*$`

	return regexp.MustCompile(r)
}

// Load templates from the tmpl directory
func loadTemplates() *template.Template {
	funcMap := template.FuncMap{
		// Any custom template functions go here
	}

	tmplFS := os.DirFS("./tmpl")

	tmpls, err := template.New("funcs").Funcs(funcMap).
		ParseFS(tmplFS, "*.html")
	return template.Must(tmpls, err)
}

// Validate the path of a request
func validatePath(path string) ([]string, error) {
	m := validPath.FindStringSubmatch(path)
	if m == nil {
		return nil, errors.New("Invalid path")
	}

	fmt.Printf("m: %v\n", m)

	return m, nil
}

// Create a handler with URL path validation
// 
// Renders a template with the same name as the base path
func makeHandlerWithTemplate(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m, err := validatePath(r.URL.Path)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r, m[1])
	}
}

// Create a basic handler with URL path validation
func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m, err := validatePath(r.URL.Path)
		_ = m
		if err != nil {
			http.NotFound(w, r)
			return
		}

		// Make sure m has valid subpaths
		// TODO: Implement subpath validation

		fn(w, r)
	}
}

// Render a template with data
func renderTemplate(w http.ResponseWriter, tmpl string, data any) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Handle the default path
func handleDefault(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		// Redirect "/" to "/home"
		http.Redirect(w, r, "/home", http.StatusFound)
	} else {
		// This url was not matched by any handlers
		http.NotFound(w, r)
		return
	}
}

// Handle the home page
func handleIndex(w http.ResponseWriter, r *http.Request, tmpl string) {
	renderTemplate(w, tmpl+".html", nil)
}

// Handle the intervals page
func handleGetIntervals(w http.ResponseWriter, r *http.Request, tmpl string) {
	// Create initial fields for inputs
	inputs := []InputField{
		{
			Value: "",
			Id:    1,
			Error: "",
		},
		{
			Value: "",
			Id:    2,
			Error: "",
		},
	}

	// Delete history on new page request
	history, err := r.Cookie("intervals-session")
	if err == nil {
		history.MaxAge = -1
		http.SetCookie(w, history)
	}

	renderTemplate(w, tmpl+".html", inputs)
}

// Handle form submissions to the intervals page
func handlePostIntervals(w http.ResponseWriter, r *http.Request) {
	p1 := r.FormValue("pitch1")
	p2 := r.FormValue("pitch2")
	o1 := r.FormValue("octave1")
	o2 := r.FormValue("octave2")

	idx1, err1 := music.Search(p1)
	idx2, err2 := music.Search(p2)
	if err1 != nil || err2 != nil {
		http.Error(w, "Invalid pitch received.", http.StatusBadRequest)
		return
	}

	octave1, _ := strconv.Atoi(o1)
	octave2, _ := strconv.Atoi(o2)

	note1 := music.Note{Pitch: idx1, Octave: octave1}
	note2 := music.Note{Pitch: idx2, Octave: octave2}

	if note1.Pitch > note2.Pitch {
		// Make the first pitch always treated
		// as an octave below the second
		note2.Octave++
	}

	distance := note1.GetInterval(note2)
	intervalName := music.IntervalToString(distance)

	buffer := [2]string{}
	buffer[0] = fmt.Sprintf("(%d) %s [%s -> %s]", distance, intervalName, p1, p2)

	// Remember the last answer
	cookie, err := r.Cookie("intervals-session")
	if err == nil {
		buffer[1] = cookie.Value
	}
	cookie = &http.Cookie{
		Name:   "intervals-session",
		Value:  buffer[0],
		MaxAge: 0,
	}
	http.SetCookie(w, cookie)

	templates.ExecuteTemplate(w, "intervalResult", buffer)
}

// Route interval handler based on request method
func handleIntervals(w http.ResponseWriter, r *http.Request, tmpl string) {
	if r.Method == http.MethodGet {
		handleGetIntervals(w, r, tmpl)
	} else if r.Method == http.MethodPost {
		handlePostIntervals(w, r)
	}
}

// Handle validation requests for the intervals page
func handleIntervalsValidation(w http.ResponseWriter, r *http.Request) {
	// Find which input this is validating
	inputName := r.Header["Hx-Trigger-Name"][0]
	inputIdentifier, err := regexp.Compile("([a-zA-Z]+)([0-9]+)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	field := inputIdentifier.FindStringSubmatch(inputName)

	inputType := strings.ToLower(field[1])
	id, err := strconv.Atoi(field[2])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	switch inputType {
	case "pitch":
		handleValidateNote(w, r, id)
		break
	case "octave":
		handleValidateOctave(w, r, id)
		break
	}
}

// Handle octave mode toggle requests for the intervals page
func handleOctaveMode(w http.ResponseWriter, r *http.Request) {
	// Replace form inputs
	switchValue := r.FormValue("advanced")
	inputs := [][]InputField{
		{
			{
				Value: "",
				Id:    1,
				Error: "",
			},
			{
				Value: "",
				Id:    2,
				Error: "",
			},
		},
		{
			{
				Value: "0",
				Id:    1,
				Error: "",
			},
			{
				Value: "0",
				Id:    2,
				Error: "",
			},
		},
	}
	// Delete history on new page request
	history, err := r.Cookie("intervals-session")
	if err == nil {
		history.MaxAge = -1
		http.SetCookie(w, history)
	}

	// Render the inputs
	if switchValue == "on" {
		renderTemplate(w, "intervalsAdvanced", inputs)
	} else {
		renderTemplate(w, "intervalsBasic", inputs[0])
	}
}

// Handle note validation requests for the intervals page
func handleValidateNote(w http.ResponseWriter, r *http.Request, id int) {
	// Validate the pitch
	inputName := fmt.Sprintf("pitch%d", id)
	pitch := strings.Trim(r.FormValue(inputName), " ")
	_, err := music.Search(pitch)
	input := &InputField{
		Id:    id,
		Value: pitch,
		Error: "",
	}
	if err != nil || len(pitch) > 4 {
		input.Error = "Invalid pitch"
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
