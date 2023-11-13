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

func handleStatic(w http.ResponseWriter, r *http.Request) {

}

func handleIndex(w http.ResponseWriter, r *http.Request, tmpl string) {
	renderTemplate(w, tmpl+".html", nil)
}

func handleIntervals(w http.ResponseWriter, r *http.Request, tmpl string) {
	if r.Method == http.MethodGet {
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

		renderTemplate(w, tmpl+".html", inputs)
	} else if r.Method == http.MethodPost {
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

		type Interval struct {
			Distance int
			String   string
		}

		interval := &Interval{}
		interval.Distance = note1.GetInterval(note2)
		interval.String = music.IntervalToString(interval.Distance)

		templates.ExecuteTemplate(w, "intervalResult", interval)
	}
}

func handleValidateNote(w http.ResponseWriter, r *http.Request) {
	inputName := r.Header["Hx-Trigger-Name"][0]
	inputIdentifier, err := regexp.Compile("[0-9]+")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	whichPitch, err := strconv.Atoi(inputIdentifier.FindString(inputName))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	pitch := strings.Trim(r.FormValue(inputName), " ")
	_, err = music.Search(pitch)
	if err != nil {
		inputError := &InputField{
			WhichPitch: whichPitch,
			Value:      pitch,
			Error:      "Invalid pitch",
		}
		templates.ExecuteTemplate(w, "invalidateNote", inputError)
	} else {
		input := &InputField{
			WhichPitch: whichPitch,
			Value:      pitch,
			Error:      "",
		}
		templates.ExecuteTemplate(w, "noteInput", input)
	}
}
