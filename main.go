package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

// Render the main index page with categories
func indexHandler(w http.ResponseWriter, r *http.Request) {
	imageDir := "./static/images/"
	files, err := ioutil.ReadDir(imageDir)
	if err != nil {
		http.Error(w, "Unable to read images directory", http.StatusInternalServerError)
		return
	}

	var categories []string
	for _, file := range files {
		if file.IsDir() {
			categories = append(categories, file.Name())
		}
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template not found: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, categories)
	if err != nil {
		http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
	}
}

// Render the category page showing items
func categoryPageHandler(w http.ResponseWriter, r *http.Request) {
	category := strings.TrimPrefix(r.URL.Path, "/category/")
	categoryPath := "./static/images/" + category

	files, err := ioutil.ReadDir(categoryPath)
	if err != nil {
		http.Error(w, "Unable to read category directory", http.StatusInternalServerError)
		return
	}

	var items []string
	for _, file := range files {
		if !file.IsDir() {
			itemName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
			items = append(items, itemName)
		}
	}

	tmpl, err := template.ParseFiles("templates/category.html")
	if err != nil {
		http.Error(w, "Template not found: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Category string
		Items    []string
	}{
		Category: category,
		Items:    items,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
	}
}

// Render item page for specific images
func itemPageHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/item/"), "/")
	if len(parts) < 2 {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}
	category := parts[0]
	item := parts[1]
	itemPath := "./static/images/" + category + "/" + item

	// Read files in the item's folder
	files, err := ioutil.ReadDir(itemPath)
	if err != nil {
		http.Error(w, "Unable to read item directory", http.StatusInternalServerError)
		return
	}

	// Collect all images except for the thumbnail
	var images []string
	for _, file := range files {
		if !file.IsDir() && file.Name() != "thumbnail.jpg" && file.Name() != "description.txt" {
			images = append(images, file.Name())
		}
	}

	// Read the description from description.txt
	descriptionPath := itemPath + "/description.txt"
	descriptionBytes, err := ioutil.ReadFile(descriptionPath)
	description := "No description available."
	if err == nil {
		description = string(descriptionBytes)
	} else {
		log.Printf("Warning: Unable to read description.txt for %s: %v", itemPath, err)
	}

	// Prepare data to send to the template
	tmpl, err := template.ParseFiles("templates/item.html")
	if err != nil {
		http.Error(w, "Template not found: "+err.Error(), http.StatusInternalServerError)
		return
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

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
	}
}




func main() {
	// Serve the main page
	http.HandleFunc("/", indexHandler)

	// Serve category pages
	http.HandleFunc("/category/", categoryPageHandler)

	// Serve item pages
	http.HandleFunc("/item/", itemPageHandler)

	// Serve static files (like images and css)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start the server
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
