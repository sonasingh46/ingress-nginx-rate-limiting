package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

/*
Build : docker build -t sonasingh46/ping .
Run: docker run -d --publish 8080:8080 sonasingh46/ping
*/
func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	mux := httprouter.New()
	mux.GET("/ping", ping)
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	log.Infof("Serving on %s ...", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to start the server: %s", err)
	}
}

func ping(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Println("Ping received...")
	fmt.Fprint(rw, "Pong!\n")
}
