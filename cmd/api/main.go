package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/enricozb/clip/lib/clip"
)

func setClip(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("read all: %v", err), http.StatusInternalServerError)
	}

	clip.Sync(body)
}

func handleRequests() {
	http.HandleFunc("/", setClip)
	log.Fatal(http.ListenAndServe(":6719", nil))
}

func main() {
	handleRequests()
}
