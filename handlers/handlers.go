package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"
)

var (
	templates = template.Must(template.ParseFiles(
		"index.html",
	))
)

func init() {
	initMarkdownFlags()
	http.HandleFunc("/render", handle)
	http.HandleFunc("/", mainHandle)
}

func mainHandle(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
	return
}

type BabelmarkResponder struct {
	Name    string `json:"name"`
	Html    string `json:"html"`
	Version string `json:"version"`
}

var defaultResponder = NewBabelmarkResponder()

func NewBabelmarkResponder() (r *BabelmarkResponder) {
	r = &BabelmarkResponder{
		Name:    "Blackfriday",
		Version: "Qortex",
	}
	return
}

func (this *BabelmarkResponder) Responde(w http.ResponseWriter, text string) {
	this.Html = renderMD(text)
	body, err := json.Marshal(this)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Write(body)
}

func handle(w http.ResponseWriter, r *http.Request) {
	input := r.FormValue("text")
	if len(input) > 1000 {
		w.WriteHeader(500)
		return
	}

	defaultResponder.Responde(w, input)
}
