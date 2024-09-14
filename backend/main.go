package main

import (
	"fmt"
	"net/http"
)

const port = ":9090"
var counter int

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /stream", streamDataHandler)
	fmt.Println("Server is running on port", port)
	http.ListenAndServe(port, mux)
}

func streamDataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// simulate data changes with counter
	counter++
	fmt.Fprintf(w, "data: %v\n\n", counter)
	w.(http.Flusher).Flush()
	r.Context().Done()

}
