package handlers

import (
	"net/http"
	"power4web/src/server/pages"
)

func LandingPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		pages.LoadPage(w, r, "src/client/404.html")
		return
	}

	pages.LoadPage(w, r, "src/client/index.html")
}
