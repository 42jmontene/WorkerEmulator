package main

import (
	"log"
	"net/http"
	"fmt"
)

func totoHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\t| Ok\n", r.URL.String())
}

func main () {
	http.HandleFunc("/run/", totoHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))	// Port d'écoute
}
