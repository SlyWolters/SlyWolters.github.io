package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

// Helper function to list directories
func listDirectories(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var dirs []string
	for _, file := range files {
		if file.IsDir() {
			dirs = append(dirs, file.Name())
		}
	}
	return dirs, nil
}

// Helper function to list files in a directory (excluding directories)
func listFiles(path string, exclude []string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var items []string
	for _, file := range files {
		if !file.IsDir() {
			name := file.Name()
			if !contains(exclude, name) {
				items = append(items, strings.TrimSuffix(name, filepath.Ext(name)))
			}
		}
	}
	return items, nil
}

// Helper function to check if a slice contains a value
func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// Generic function to render templates
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplParsed, err := template.ParseFiles("templates/layout.html", "templates/"+tmpl+".html")
	if err != nil {
		http.Error(w, "Template not found: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmplParsed.Execute(w, data)
	if err != nil {
		http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
	}
}

// Index handler - Shows categories
func indexHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := listDirectories("./static/images/")
	if err != nil {
		http.Error(w, "Unable to read images directory", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title      string
		Categories []string
	}{
		Title:      "Photo Gallery",
		Categories: categories,
	}

	renderTemplate(w, "index", data)
}

// Category handler - Shows items in a category
func categoryPageHandler(w http.ResponseWriter, r *http.Request) {
	category := strings.TrimPrefix(r.URL.Path, "/category/")
	items, err := listFiles("./static/images/"+category, nil)
	if err != nil {
		http.Error(w, "Unable to read category directory", http.StatusInternalServerError)
		return
	}

	data := struct {
		Category string
		Items    []string
	}{
		Category: category,
		Items:    items,
	}

	renderTemplate(w, "category", data)
}

// Item handler - Shows details for an item
func itemPageHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/item/"), "/")
	if len(parts) < 2 {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}

	category, item := parts[0], parts[1]
	itemPath := "./static/images/" + category + "/" + item

	images, err := listFiles(itemPath, []string{"thumbnail.jpg", "description.txt"})
	if err != nil {
		http.Error(w, "Unable to read item directory", http.StatusInternalServerError)
		return
	}

	// Read description
	descriptionPath := itemPath + "/description.txt"
	descriptionBytes, err := ioutil.ReadFile(descriptionPath)
	description := "No description available."
	if err == nil {
		description = string(descriptionBytes)
	} else {
		log.Printf("Warning: Unable to read description.txt for %s: %v", itemPath, err)
	}

	data := struct {
		Category    string
		Item        string
		Description string
		Images      []string
	}{
		Category:    category,
		Item:        item,
		Description: description,
		Images:      images,
	}

	renderTemplate(w, "item", data)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/category/", categoryPageHandler)
	http.HandleFunc("/item/", itemPageHandler)

	// Serve static files (like images and css)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server started on :8181")
	http.ListenAndServe(":8181", nil)
}
