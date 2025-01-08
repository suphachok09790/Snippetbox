package main

import (
	"path/filepath"
	"html/template"
	"snippetbox.suphachok.net/internal/models"
)

type templateData struct {
	Snippet models.Snippet
	Snippets []models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	// Initialize a new map to act as the cache.
	cache := map[string]*template.Template{}
	
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
	return nil, err
	}
	// Loop through the page filepaths one-by-one.
	for _, page := range pages {
	// Extract the file name (like 'home.tmpl') from the full filepath
	// and assign it to the name variable.
	name := filepath.Base(page)

	ts, err := template.ParseFiles("./ui/html/base.tmpl")
	if err != nil {
		return nil, err
	}

	ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")
	if err != nil {
		return nil, err
	}

	ts, err = ts.ParseFiles(page)
	if err != nil {
		return nil, err
	}

	cache[name] =ts
	}
	// Return the map.
	return cache, nil
}
	

