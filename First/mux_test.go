package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/Zero", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Haii Zero...")
	})
	mux.HandleFunc("/Owen", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Haii Owen...")
	})
	mux.HandleFunc("/Example/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Contoh 1")
	})
	mux.HandleFunc("/Example/testing/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Contoh 2")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
