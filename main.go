package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kr0106686/public/filer"
)

func main() {
	socketURL := os.Getenv("SOCKET_URL")
	port := os.Getenv("FRONT_PORT")
	staticDir := "/static"

	log.Println(socketURL)
	log.Println(port)
	log.Println(staticDir)

	f := filer.New()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f.ReadHtmlTmpl(w, "home", map[string]interface{}{
			"title": "home",
		})
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		f.ReadHtmlTmpl(w, "login", map[string]interface{}{
			"title": "login",
		})
	})

	http.HandleFunc("/omok", func(w http.ResponseWriter, r *http.Request) {
		f.ReadHtmlTmpl(w, "login", map[string]interface{}{
			"title":     "omok",
			"socketURL": socketURL,
		})
	})

	log.Printf("http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
