package pages

import (
	"html/template"
	"net/http"
	"path/filepath"
	"power4web/src/server/data"
)

func LoadPage(w http.ResponseWriter, r *http.Request, pagePath string) {
	funcs := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"minus": func(a, b int) int {
			return a - b
		},
	}

	base := filepath.Base(pagePath)

	tmpl, err := template.New(base).Funcs(funcs).ParseFiles(pagePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, base, data.ServerData)
}
