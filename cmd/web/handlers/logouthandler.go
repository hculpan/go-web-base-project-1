package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/hculpan/go-web-base-project-1/cmd/web/templates"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	})

	comp := templates.LogoutTemplate(appTitle)
	comp.Render(context.Background(), w)
}
