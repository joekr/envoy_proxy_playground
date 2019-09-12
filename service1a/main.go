package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"os"
)

type serviceHandler struct {
}

func (h *serviceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/service/")
	requestID := r.Header.Get("x-request-id")
	canary := r.Header.Get("x-canary-version")
	hostname, _ := os.Hostname()
	log.Println("x-request-id ", requestID, r.URL.Path)
	log.Println("x-canary-version ", canary)
	log.Println("service id ", id)

	w.WriteHeader(500)
	fmt.Fprintf(w, "ERROR from behind Envoy service1a Host: %s\n", hostname)
}

func main() {
	f, err := os.OpenFile("service.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
			log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	http.Handle("/service/", new(serviceHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
