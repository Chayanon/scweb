package view

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var (
	tpIndex = template.New("")
)

func init() {
	fmt.Println("TempInit")
	tpIndex.Funcs(template.FuncMap{})
	_, err := tpIndex.ParseFiles("template/root.tmpl", "template/index.tmpl")
	if err != nil {
		panic(err)
	}
	tpIndex = tpIndex.Lookup("root")
}

func render(t *template.Template, w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.Execute(w, data)
	if err != nil {
		log.Println(err)
		fmt.Println("TT")
	}
}

func Index(w http.ResponseWriter, data interface{}) {
	render(tpIndex, w, data)
}

// var (
// 	tpIndex = parseTemplate("root.tmpl", "index.tmpl")
// )

// var m = minify.New()

// const templateDir = "template"

// func init() {
// 	m.AddFunc("text/html", html.Minify)
// 	m.AddFunc("text/css", css.Minify)
// 	m.AddFunc("text/javascript", js.Minify)
// }

// func joinTemplateDir(files ...string) []string {
// 	r := make([]string, len(files))
// 	for i, f := range files {
// 		r[i] = filepath.Join(templateDir, f)
// 	}
// 	return r
// }

// func parseTemplate(files ...string) *template.Template {
// 	t := template.New("")
// 	t.Funcs(template.FuncMap{})
// 	_, err := t.ParseFiles(joinTemplateDir(files...)...)
// 	if err != nil {
// 		panic(err)
// 	}
// 	t = t.Lookup("root")
// 	return t
// }

// func render(t *template.Template, w http.ResponseWriter, data interface{}) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	buf := bytes.Buffer{}
// 	err := t.Execute(&buf, data)
// 	// t.Excute => buf => m.Minify => w
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	m.Minify("text/html", w, &buf)
// }
