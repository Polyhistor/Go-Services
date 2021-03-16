package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

// helper that goes into the templates directory and renders the corresponding HTML template
func RenderTemplate (w http.ResponseWriter, tmpl string) {
	_, err := RenderTemplateTest(w)

	if err != nil {
		fmt.Println("error getting template from the templace cache", err)
	}

	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err = parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing template")
	}
}

var functions = template.FuncMap{}

func RenderTemplateTest(w http.ResponseWriter) (map[string]*template.Template, error) {
	// creating a map
	myCache := map[string]*template.Template{}

	fmt.Println("mycache:", myCache)

	// getting all the pages from the templates folder
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	fmt.Println("pages:", pages)

	// handling err for pages
	if err != nil {
		return myCache, err
	}

	// iterating through the collection
	for _, page := range pages {
		// getting the name of the file not the full path
		name := filepath.Base(page)

		fmt.Println("name:", name)

		// creating a new template from the name, adding custom functions and then parse it
		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)

		// checking if there is any error in template creation
		if err != nil {
			return myCache, err
		}

		// getting all the layouts
		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		// checking for errors for the layout
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")

			if err != nil {
				return myCache, err
			}

			myCache[name] = templateSet
			return myCache, nil
		}
	}
	return myCache, nil
}
