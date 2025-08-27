package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// PageData holds the data to be passed to templates
type PageData struct {
	Title   string
	Content template.HTML
}

var templates map[string]*template.Template

// LoadTemplates initializes all templates with header and footer
func LoadTemplates() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	// List of page templates
	pages := []string{
		"home.html",
		"about.html",
		"contact.html",
		"skills.html",
		"education.html",
		"experience.html",
	}

	// Load each template with the common header and footer
	for _, page := range pages {
		files := []string{
			filepath.Join("templates", "header.html"),
			filepath.Join("templates", page),
			filepath.Join("templates", "footer.html"),
		}

		templates[page] = template.Must(template.ParseFiles(files...))
	}
}

// RenderTemplate renders a template with the given data
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) error {
	// If templates are not loaded yet, load them
	if templates == nil {
		LoadTemplates()
	}

	// Get the template from the cache
	t, ok := templates[tmpl]
	if !ok {
		return http.ErrNotSupported
	}

	// Execute the template
	return t.ExecuteTemplate(w, tmpl, data)
}
