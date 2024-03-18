package handlers

import (
	"context"
	"net/http"

	"github.com/hculpan/go-web-base-project-1/cmd/web/templates"
)

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	comp := templates.InfoTemplate(appTitle)
	comp.Render(context.Background(), w)
}
