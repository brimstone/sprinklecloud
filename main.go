package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	apachelog "github.com/lestrrat-go/apache-logformat"
)

type fileHandler struct{}

func (h *fileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hostparts := strings.Split(r.Host, ":")
	path := path.Join("www", hostparts[0])
	if strings.HasPrefix(path, "www") {
		headers, err := os.Open(path + "/" + r.URL.Path + ".headers")
		if err == nil {
			defer headers.Close()

			scanner := bufio.NewScanner(headers)
			for scanner.Scan() {
				header := strings.SplitN(scanner.Text(), ":", 2)
				w.Header().Set(header[0], header[1])
			}
		}
		http.FileServer(http.Dir(path)).ServeHTTP(w, r)
		return
	}
	http.Error(w, "Client error in request", 400)

}

func main() {
	l, _ := apachelog.New(`%h %l %u %t "%r" %>s %b "%{Referer}i" "%{User-agent}i" "%v"`)
	s := &http.Server{
		Addr:           ":8080",
		Handler:        l.Wrap(&fileHandler{}, os.Stderr),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Starting server")
	log.Fatal(s.ListenAndServe())
}
