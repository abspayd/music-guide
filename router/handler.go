package router

import (
	"html/template"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/abspayd/music-companion/music"
)

var (
	templates = loadTemplates()
)

type InputField struct {
	Id    int
	Value string
	Error string
}

type IntervalResult struct {
	IntervalName string `json:"interval_name"`
	Distance     int    `json:"distance"`
	Note1        string `json:"note1"`
	Note2        string `json:"note2"`
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

		fn(w, r, m[0])
	}
}

// Create a basic handler with URL path validation
func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := validatePath(r.URL.Path)
		if err != nil {
			http.NotFound(w, r)
			return
		}

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

	renderTemplate(w, tmpl+".html", inputs)
}

// Handle form submissions to the intervals page
func handlePostIntervals(w http.ResponseWriter, r *http.Request) {
	n1 := r.FormValue("note1")
	n2 := r.FormValue("note2")
	o1 := r.FormValue("octave1")
	o2 := r.FormValue("octave2")

	idx1, err1 := music.Search(n1)
	idx2, err2 := music.Search(n2)
	if err1 != nil || err2 != nil {
		http.Error(w, "Invalid note received.", http.StatusBadRequest)
		return
	}

	octave1, _ := strconv.Atoi(o1)
	octave2, _ := strconv.Atoi(o2)

	pitch1 := music.Pitch{Note: idx1, Octave: octave1}
	pitch2 := music.Pitch{Note: idx2, Octave: octave2}

	if octave1 == octave2 && pitch1.Note > pitch2.Note {
		// Make the first pitch always treated
		// as an octave below the second
		pitch2.Octave++
	}

	distance := pitch1.GetInterval(pitch2)
	intervalName := music.IntervalToString(distance)

	result := IntervalResult{
		// Current result
		IntervalName: intervalName,
		Distance:     distance,
		Note1:        n1,
		Note2:        n2,
	}

	templates.ExecuteTemplate(w, "intervalResult", result)
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
	case "note":
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
	history, err := r.Cookie("interval-between")
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
