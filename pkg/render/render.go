package render

import (
	"bytes"
	"github.com/polyhistor/go-service/pkg/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplate(a *config.AppConfig) {
	app = a
}

// helper that goes into the templates directory and renders the corresponding HTML template
func RenderTemplate (w http.ResponseWriter, tmpl string) {
	// get the template cache from the app config
	templateCache := app.TemplateCache

	template, doesTemplateExist := templateCache[tmpl]

	if !doesTemplateExist {
		log.Fatal("could not get template from template cache")
	}

	buffer := new(bytes.Buffer)
	_ = template.Execute(buffer, nil)
	_, err := buffer.WriteTo(w)

	if err != nil {
		log.Fatal("error writing template to browser")
	}

}

var functions = template.FuncMap{}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	// creating a map
	myCache := map[string]*template.Template{}

	// getting all the pages from the templates folder
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	// handling err for pages
	if err != nil {
		return myCache, err
	}

	// iterating through the collection
	for _, page := range pages {
		// getting the name of the file not the full path
		name := filepath.Base(page)
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
		}

		myCache[name] = templateSet
	}
	return myCache, nil
}
