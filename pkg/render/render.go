package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

// helper that goes into the templates directory and renders the corresponding HTML template
func RenderTemplate (w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template")
	}
}


func RenderTemplateTest(w http.ResponseWriter, tmpl string) error {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err !=nil {
		return err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name)
	}
}