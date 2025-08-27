package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

// DEVELOPMENT MODE TOGGLE
// Set to true during development to reload templates on every request
// Set to false for production to use template caching
const DEVELOPMENT_MODE = true

// Template for our pages
var tmpl *template.Template

// Data to pass to templates
type TemplateData struct {
	Title string
	Year  int // Current year for copyright notices
}

func init() {
	// Parse all templates at once for initial load
	loadTemplates()
}

// Load templates - called on init and optionally on each request in dev mode
func loadTemplates() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
	log.Println("Templates loaded")
}

// Handler functions for different routes
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

// Render a template with development mode support
func renderTemplate(w http.ResponseWriter, tmplName string, title string) {
	// In development mode, reload templates on each request
	if DEVELOPMENT_MODE {
		loadTemplates()
	}

	// Create template data with current year
	data := TemplateData{
		Title: title,
		Year:  time.Now().Year(),
	}

	// Execute template with data
	err := tmpl.ExecuteTemplate(w, tmplName, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	// Serve static files properly with path stripping
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Route handlers for all pages
	http.HandleFunc("/", home)
	http.HandleFunc("/about.html", about)
	http.HandleFunc("/contact.html", contact)
	http.HandleFunc("/skills.html", skills)

	// Start server and listen
	log.Println("Server starting on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
