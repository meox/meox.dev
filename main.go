package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/meox/meox.dev/contentwriter"

	"github.com/meox/meox.dev/cv_parser"
	log "github.com/sirupsen/logrus"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", myCv)

	log.Infof("server starting on port %s", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
	log.Info("server closed: %v", err)
}

func myCv(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		log.Warn("bath method: %s", req.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// instantiate the right content-writer
	var contentWriter contentwriter.ContentWriter
	accept := req.Header.Get("Accept")
	if strings.Contains(accept, "text/html") {
		contentWriter = contentwriter.NewHtml(w)
	} else {
		contentWriter = contentwriter.NewPlain(w)
	}

	contentWriter.ContentType()

	fd, err := cv_parser.Open("cv/meox.md")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Errorf("error opening cv: %v", err)
		return
	}

	err = contentWriter.Header()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorf("cannot write header: %v", err)
		return
	}

	body, err := fd.Parse(contentWriter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorf("error parsing cv: %v", err)
		return
	}

	_, err = w.Write([]byte(body))
	if err != nil {
		log.Warnf("writing resp to the server: %v", err)
	}
}
