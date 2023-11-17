package router

import (
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
	validPath = regexp.MustCompile("^/(home|intervals)$")
	templates = loadTemplates()
)

type InputField struct {
	Id    int
	Value string
	Error string
}

func loadTemplates() *template.Template {
	funcMap := template.FuncMap{
		"add": func(x int, y int) int {
			return x + y
		},
	}

	tmplFS := os.DirFS("./tmpl")

	tmpls, err := template.New("funcs").Funcs(funcMap).
		ParseFS(tmplFS, "*.html")
	return template.Must(tmpls, err)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)

		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[1])
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, data any) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

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

func handleIndex(w http.ResponseWriter, r *http.Request, tmpl string) {
	renderTemplate(w, tmpl+".html", nil)
}

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

func handleIntervals(w http.ResponseWriter, r *http.Request, tmpl string) {
	if r.Method == http.MethodGet {
		handleGetIntervals(w, r, tmpl)
	} else if r.Method == http.MethodPost {
		handlePostIntervals(w, r)
	}
}

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
