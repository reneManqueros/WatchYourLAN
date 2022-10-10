package main

import (
	"html/template"
	"net/http"
	"watchyourlan/models"
)

func offline(w http.ResponseWriter, r *http.Request) {
	type allData struct {
		Config models.Conf
		Hosts  []models.Host
	}
	var guiData allData
	guiData.Config = models.AppConfig
	guiData.Hosts = []models.Host{}

	for _, oneHost := range AllHosts {
		if oneHost.Now == 0 {
			guiData.Hosts = append(guiData.Hosts, oneHost)
		}
	}

	tmpl, _ := template.ParseFiles("templates/offline.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "offline", guiData)
}

func online(w http.ResponseWriter, r *http.Request) {
	type allData struct {
		Config models.Conf
		Hosts  []models.Host
	}
	var guiData allData
	guiData.Config = models.AppConfig
	guiData.Hosts = []models.Host{}

	for _, oneHost := range AllHosts {
		if oneHost.Now == 1 {
			guiData.Hosts = append(guiData.Hosts, oneHost)
		}
	}

	tmpl, _ := template.ParseFiles("templates/offline.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "offline", guiData)
}
