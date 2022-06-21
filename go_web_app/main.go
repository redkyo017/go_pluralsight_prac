package main

import (
	"net/http"
)

func main() {
	// log.Println("con co be be")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ha ha ha ha"))
	})
	http.ListenAndServe(":8000", nil)
}
