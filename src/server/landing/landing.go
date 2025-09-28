package landing

import (
	"html/template"
	"net/http"
	"power4web/src/server/data"
)

func loadPage(w http.ResponseWriter, r *http.Request) {
	funcs := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
	}

	tmpl, err := template.New("index.html").Funcs(funcs).ParseFiles("src/client/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "index.html", data.ServerData)
}

func LoadLanding(w http.ResponseWriter, r *http.Request) {
	loadPage(w, r)
}
