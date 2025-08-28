package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

// DEVELOPMENT MODE TOGGLE
const DEVELOPMENT_MODE = false

var tmpl *template.Template

type TemplateData struct {
	Title string
	Year  int
}

func init() {
	loadTemplates()
}

func loadTemplates() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
	log.Println("Templates loaded")
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	renderTemplate(w, "home.html", "Home")
}

func about(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html", "About Me")
}

func contact(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact.html", "Contact Me")
}

func skills(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "skills.html", "My Skills")
}

func renderTemplate(w http.ResponseWriter, tmplName string, title string) {
	if DEVELOPMENT_MODE {
		loadTemplates()
	}

	data := TemplateData{
		Title: title,
		Year:  time.Now().Year(),
	}

	err := tmpl.ExecuteTemplate(w, tmplName, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Routes
	http.HandleFunc("/", home)
	http.HandleFunc("/about.html", about)
	http.HandleFunc("/contact.html", contact)
	http.HandleFunc("/skills.html", skills)

	// ✅ Use Render's PORT env variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local dev
	}

	log.Printf("Server starting on port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
