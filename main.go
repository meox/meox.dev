package main

import (
	"fmt"
	"net/http"

	"github.com/meox/meox.dev/cv_parser"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

func main() {
	noSSL := pflag.Bool("no-ssl", false, "non SSL mode")
	certFile := pflag.String("cert-file", "", "cert file")
	keyFile := pflag.String("key-file", "", "key file")
	address := pflag.String("listen-ip", "0.0.0.0", "listen address")
	port := pflag.Int("port", 3000, "server port")

	pflag.Parse()

	if !*noSSL {
		if certFile == nil {
			log.Fatalf("missing cert file")
		}
		if keyFile == nil {
			log.Fatalf("missing key file")
		}
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", myCv)

	log.Infof("server starting on port %d", *port)

	var err error
	if *noSSL {
		err = http.ListenAndServe(fmt.Sprintf("%s:%d", *address, *port), mux)
	} else {
		err = http.ListenAndServeTLS(fmt.Sprintf("%s:%d", *address, *port), *certFile, *keyFile, mux)
	}

	if err != nil {
		log.Fatalf("server exited: %v", err)
	}
	log.Info("server closed")
}

func myCv(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		log.Warn("bath method: %s", req.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "plain/text; charset=utf-8")

	fd, err := cv_parser.Open("cv/meox.md")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Errorf("error opening cv: %v", err)
		return
	}

	body, err := fd.Parse()
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
