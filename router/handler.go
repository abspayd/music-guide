package router

import (
	"html/template"
	"net/http"
	"os"
	"regexp"

	"github.com/abspayd/music-companion/music"
)

var (
	validPath = regexp.MustCompile("^/(home|intervals)$")
	templates = loadTemplates()
)

func loadTemplates() *template.Template{
	fs := os.DirFS("./templates")
	return template.Must(template.ParseFS(fs, "*.html"))
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

func handleIntervals(w http.ResponseWriter, r *http.Request, tmpl string) {
	p1 := r.FormValue("pitch1")
	p2 := r.FormValue("pitch2")

	var answer string
	if len(p1) > 0 && len(p2) > 0 {
		note1 := music.Note{Pitch: music.Search(p1), Octave: 0} 
		note2 := music.Note{Pitch: music.Search(p2), Octave: 0} 

		interval := note1.GetInterval(note2)
		answer = music.IntervalToString(interval)
	}
	renderTemplate(w, tmpl+".html", answer)
}

func handleGetInterval(w http.ResponseWriter, r *http.Request) {
	p1 := r.FormValue("pitch1")
	p2 := r.FormValue("pitch2")

	note1 := music.Note{Pitch: music.Search(p1), Octave: 0} 
	note2 := music.Note{Pitch: music.Search(p2), Octave: 0} 

	interval := note1.GetInterval(note2)
	intervalString := music.IntervalToString(interval)
	
	renderTemplate(w, "intervals.html", intervalString)
}

