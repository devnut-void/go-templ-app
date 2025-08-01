package handlers

import (
	"go-templ-app/views"
	"go-templ-app/views/components"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	views.Index().Render(r.Context(), w)
}

func HandleGreet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		name = "dunia"
	}

	components.Greeting(name).Render(r.Context(), w)
}
