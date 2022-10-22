package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"watchyourlan/models"
)

func index(w http.ResponseWriter, r *http.Request) {
	type allData struct {
		Config models.Conf
		Hosts  []models.Host
	}
	var guiData allData
	guiData.Config = models.AppConfig
	guiData.Hosts = AllHosts

	tmpl, _ := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	tmpl.ExecuteTemplate(w, "index", guiData)
}

func update_host(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	name := r.FormValue("name")
	knownStr := r.FormValue("known")

	if id == "" {
		fmt.Fprintf(w, "No data!")
	} else {
		var known uint16
		known = 0
		if knownStr == "on" {
			known = 1
		}

		for i, oneHost := range AllHosts {
			if oneHost.Id == id {
				AllHosts[i].Name = name
				AllHosts[i].Known = known
				AllHosts[i].Update()
			}
		}
	}

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if models.AppConfig.GuiAuth == "" {
			next.ServeHTTP(w, r)
			return
		}

		username, password, ok := r.BasicAuth()
		if ok {
			userCredentials := fmt.Sprintf(`%s:%s`, username, password)
			if userCredentials == models.AppConfig.GuiAuth {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

func webgui() {
	// fmt.Println(FoundHosts)
	address := models.AppConfig.GuiIP + ":" + models.AppConfig.GuiPort

	log.Println("=================================== ")
	log.Println(fmt.Sprintf("Web GUI at http://%s", address))
	log.Println("=================================== ")

	http.HandleFunc("/", basicAuth(index))
	http.HandleFunc("/home/", basicAuth(home))
	http.HandleFunc("/offline/", basicAuth(offline))
	http.HandleFunc("/online/", basicAuth(online))
	http.HandleFunc("/search_hosts/", basicAuth(search_hosts))
	http.HandleFunc("/sort_hosts/", basicAuth(sort_hosts))
	http.HandleFunc("/theme/", basicAuth(theme))
	http.HandleFunc("/update_host/", basicAuth(update_host))
	http.ListenAndServe(address, nil)
}
