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
	WhichPitch int
	Value      string
	Error      string
}

func loadTemplates() *template.Template {
	tmplFS := os.DirFS("./tmpl")
	return template.Must(template.ParseFS(tmplFS, "*.html"))
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
			Value:      "",
			WhichPitch: 1,
			Error:      "",
		},
		{
			Value:      "",
			WhichPitch: 2,
			Error:      "",
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

	idx1, err1 := music.Search(p1)
	idx2, err2 := music.Search(p2)
	if err1 != nil || err2 != nil {
		http.Error(w, "Invalid pitch received.", http.StatusBadRequest)
		return
	}

	note1 := music.Note{Pitch: idx1, Octave: 0}
	note2 := music.Note{Pitch: idx2, Octave: 0}

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
		Name:  "intervals-session",
		Value: buffer[0],
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

func handleValidateNote(w http.ResponseWriter, r *http.Request) {
	// Find which input this is validating
	inputName := r.Header["Hx-Trigger-Name"][0]
	inputIdentifier, err := regexp.Compile("[0-9]+")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	whichPitch, err := strconv.Atoi(inputIdentifier.FindString(inputName))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Validate the pitch
	pitch := strings.Trim(r.FormValue(inputName), " ")
	_, err = music.Search(pitch)
	input := &InputField{
		WhichPitch: whichPitch,
		Value:      pitch,
		Error:      "",
	}
	if err != nil || len(pitch) > 4 {
		input.Error = "Invalid pitch"
		templates.ExecuteTemplate(w, "invalidNote", input)
	} else {
		templates.ExecuteTemplate(w, "noteInput", input)
	}
}
