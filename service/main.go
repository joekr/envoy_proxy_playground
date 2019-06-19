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
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, " Hello from behind Envoy service%s Host: %s\n", id, hostname)
}

func main() {
	http.Handle("/service/", new(serviceHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
