package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
)

// set up context wireframe to use in templates
type templateData struct {
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float32
	Data map[string]interface{}
	CSRFToken string
	Flash string
	Warning string
	Error string
	IsAuthenticated int
	API string
	CSSVersion string
}

// set up var to hold functions that I may want to add in the future
var functions = template.FuncMap{

}

//go:embed templates
// embed allows to compile the application including all associated templates into a single binary
var templateFS embed.FS

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	return td
}

func (app *application) renderTemplate(w http.ResponseWriter, r *http.Request, page string, td *templateData, partials ...string) error {
	var t *template.Template
	var err error

	templateToRender := fmt.Sprintf("templates/%s.page.gohtml", page) // generates expected path for templates

	_, templateInMap := app.templateCache[templateToRender] // templateInMap will be true if template is found in cache

	if app.config.env == "production" && templateInMap {
		t = app.templateCache[templateToRender]
	} else {

	}
}