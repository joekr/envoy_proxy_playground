package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type serviceHandler struct {
}

func (h *serviceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/service/")

	fmt.Fprintf(w, " Hello from behind Envoy %d\n", id)
}

func main() {
	http.Handle("/service/", new(serviceHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
