package main

import (
	"flag"
	"log"
	"net/http"
	"path"
	"text/template"
)

var staticDir = flag.String("static_dir", "static", "스태틱 폴더 경로")
var socketURL = flag.String("socket_url", "ws://localhost:5555/ws", "소켓 경로")
var port = flag.String("port", "5555", "포트")

func renderHTML(layoutTemplatePath string) *template.Template {
	return template.Must(template.ParseGlob(layoutTemplatePath))
}

func main() {
	flag.Parse()

	layoutTemplatePath := path.Join(*staticDir, "templates", "layout", "*.html")
	pageTemplateDir := path.Join(*staticDir, "templates", "page")

	layoutTmpl := renderHTML(layoutTemplatePath)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(*staticDir))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		layoutTmpl.ParseFiles(path.Join(pageTemplateDir, "home.html"))
		layoutTmpl.ExecuteTemplate(w, "layout", map[string]interface{}{
			"title": "home",
		})
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		layoutTmpl.ParseFiles(path.Join(pageTemplateDir, "login.html"))
		layoutTmpl.ExecuteTemplate(w, "layout", map[string]interface{}{
			"title": "login",
		})
	})

	http.HandleFunc("/omok", func(w http.ResponseWriter, r *http.Request) {
		layoutTmpl.ParseFiles(path.Join(pageTemplateDir, "omok.html"))
		layoutTmpl.ExecuteTemplate(w, "layout", map[string]interface{}{
			"title":     "omok",
			"socketURL": *socketURL,
		})
	})

	log.Printf("http://localhost:%s", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
