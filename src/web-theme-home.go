package main

import (
	"html"
	"net/http"
	"strings"
	"watchyourlan/models"
)

func theme(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		urlString := html.EscapeString(r.URL.Path)
		tags := strings.Split(urlString, "/")
		oneTheme := tags[2]

		models.AppConfig.Theme = oneTheme
		models.AppConfig.Set()
	}

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func home(w http.ResponseWriter, r *http.Request) {
	AllHosts = models.HostsGetAll()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
