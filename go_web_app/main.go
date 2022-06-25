package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

// type myHandler struct {
// 	greating string
// }

// func (myHl myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte(fmt.Sprintf("%v world", myHl.greating)))
// }

func main() {
	// log.Println("con co be be")
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("ha ha ha ha"))
	// })
	// http.Handle("/", &myHandler{greating: "bonjour"})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// http.ServeFile(w, r, "public"+r.URL.Path)
		f, err := os.Open("public" + r.URL.Path)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
		defer f.Close()
		// var contentType string
		// switch {
		// case strings.HasSuffix(r.URL.Path, "html"):
		// 	contentType = "text/html"
		// case strings.HasSuffix(r.URL.Path, "css"):
		// 	contentType = "text/css"
		// case strings.HasSuffix(r.URL.Path, "png"):
		// 	contentType = "image/png"
		// default:
		// 	contentType = "text/plain"
		// }
		// w.Header().Add("Content-Type", contentType)
		io.Copy(w, f)
	})
	http.ListenAndServe(":8000", nil)
}
