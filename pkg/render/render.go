package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// helper that goes into the templates directory and renders the corresponding HTML template
func RenderTemplate (w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template")
	}
}
