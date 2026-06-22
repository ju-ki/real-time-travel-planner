package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"real-time-travel-planner/config"
	"real-time-travel-planner/utils"
	"regexp"
	"text/template"
)

type Page struct {
	Title string
	Body  []byte
}

var validPath = regexp.MustCompile("^/(view)/([a-zA-Z0-9]+)$")
var templates = template.Must(template.ParseFiles("templates/view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	path, _ := filepath.Abs("./templates/")

	t, err := template.ParseFiles(path + "/" + tmpl + ".html")
	if err != nil {
		fmt.Fprintf(w, "Error parsing template: %v", err)
		return
	}

	t.Execute(w, p)
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", fmt.Errorf("invalid page title")
	}
	return m[2], nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		log.Printf("Error getting title: %v", err)
		return
	}

	renderTemplate(w, "view", &Page{Title: title})
}

func main() {
	utils.LoggingSettings(config.AppConfig.LogFile)
	log.Println("Starting server on :8080")
	// http.HandleFunc("/view/", viewHandler)
	// log.Fatal(http.ListenAndServe(":8080", nil))
}
