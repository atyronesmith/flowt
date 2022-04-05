package main

import (
	"log"
	"net/http"

	"github.com/atyronesmith/flowt/internal/graph"
	"github.com/go-echarts/examples/examples"
)

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	examplers := []examples.Exampler{
		graph.TreeExamples{},
	}

	for _, e := range examplers {
		e.Examples()
	}

	fs := http.FileServer(http.Dir("web/html"))
	log.Println("running server at http://localhost:8089")
	log.Fatal(http.ListenAndServe("localhost:8089", logRequest(fs)))
}
