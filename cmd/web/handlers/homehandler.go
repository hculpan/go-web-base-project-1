package handlers

import (
	"context"
	"net/http"

	"github.com/hculpan/go-web-base-project-1/cmd/web/templates"
	"github.com/hculpan/go-web-base-project-1/internal/auth"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if !auth.IsAuthorized(r) {
		http.Redirect(w, r, "/info", http.StatusSeeOther)
		return
	}

	comp := templates.HomeTemplate(appTitle)
	comp.Render(context.Background(), w)
}
