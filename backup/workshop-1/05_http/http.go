package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

type Handler struct{}

func main() {
	http.ListenAndServe(":8080", &Handler{})
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/somefile" {
		f, err := os.Open("somefile.txt")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			w.Write([]byte{'\n'})
			return
		}
		defer f.Close()

		w.Header().Add("Content-Type", "text/plain")
		_, err = io.Copy(w, f)
		if err != nil {
			log.Println("io.Copy:", err)
		}
		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not found\n"))
}
