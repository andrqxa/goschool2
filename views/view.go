package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	LayoutDir   string = "views/layouts/"
	TemplateDir string = "views/"
	TemplateExt string = ".gohtml"
)

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}

// Render is used to render the view with predefined layout.
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

// layoutFiles returns a clise of strings representing the layout files used in
// our applicaton.
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

func NewView(layout string, files ...string) *View {
	addTemplatePath(files)
	addTemplateExt(files)
	files = append(files, layoutFiles()...)
	templ, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: templ,
		Layout:   layout,
	}
}

// addTemplatePath takes in a slice of strings representing file path for
// templates, and it prepends the TemplateDir directory to each string in the
// slice
//
// Eg the input {"home"} would result in the output {"views/home"} if
// TemplateDir == "views/"
func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = TemplateDir + f
	}
}

// addTemplatePath takes in a slice of strings representing file path for
// templates, and it prepends the TemplateExt extension to each string in the
// slice
//
// Eg the input {"home"} would result in the output {"home.gohtml"} if
// TemplateExt == ".gohtml"
func addTemplateExt(files []string) {
	for i, f := range files {
		files[i] = f + TemplateExt
	}
}
