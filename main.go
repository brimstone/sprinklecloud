package main

import (
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	apachelog "github.com/lestrrat-go/apache-logformat"
)

type myHandler struct{}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := path.Join("www", r.Host)
	if strings.HasPrefix(path, "www") {
		http.FileServer(http.Dir(path)).ServeHTTP(w, r)
		return
	}
	http.Error(w, "Client error in request", 400)

}

func main() {
	l, _ := apachelog.New(`%h %l %u %t "%r" %>s %b "%{Referer}i" "%{User-agent}i" "%v"`)
	s := &http.Server{
		Addr:           ":8080",
		Handler:        l.Wrap(&myHandler{}, os.Stderr),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Starting server")
	log.Fatal(s.ListenAndServe())
}
