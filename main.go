package main

import (
	"io"
	"net/http"
	_ "net/http/pprof"

	"github.com/jrieck1991/runlog/internal/app"

	log "github.com/sirupsen/logrus"
)

const addr string = "localhost:6060"

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	go func() {
		serve()
	}()
	app.Do()
}

func serve() error {

	infoHandler := func(w http.ResponseWriter, req *http.Request) {
		log.SetLevel(log.InfoLevel)
		io.WriteString(w, "Set to info!\n")
	}
	http.HandleFunc("/info", infoHandler)

	debugHandler := func(w http.ResponseWriter, req *http.Request) {
		log.SetLevel(log.DebugLevel)
		io.WriteString(w, "Set to debug!\n")
	}
	http.HandleFunc("/debug", debugHandler)

	if err := http.ListenAndServe(addr, nil); err != nil {
		return err
	}

	return nil
}
