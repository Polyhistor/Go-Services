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

var functions = template.FuncMap{

}


func RenderTemplateTest(w http.ResponseWriter, tmpl string) (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}

			myCache[name] = ts
			return myCache, ts
		}

		return myCache, nil

	}
}