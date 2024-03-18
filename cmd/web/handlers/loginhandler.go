package handlers

import (
	"context"
	"net/http"

	"github.com/hculpan/go-web-base-project-1/cmd/web/templates"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	errorMsg := ""
	if errQuery := r.URL.Query().Get("error"); errQuery != "" {
		if errQuery == "invalid_credentials" {
			errorMsg = "Invalid username or password. Please try again."
		}
		// Handle other error types here as needed
	}
	// Execute the template to generate HTML content
	comp := templates.LoginTemplate(appTitle, errorMsg)
	comp.Render(context.Background(), w)
}
