package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	graph "github.com/atyronesmith/flowt/internal/graph"
	"github.com/go-echarts/examples/examples"
)

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	const (
		charFileUsage = "Name of chart file."
		charFileName = "cmd/app/fixtures/npmdepgraph.json"
	) 

	var chartFile string

	isVerbose := flag.Bool("verbose", false, "Print extra runtime information.")
	isHelp := flag.Bool("help", false, "Print usage information.")
	flag.StringVar(&chartFile, "chart", chartFile, charFileUsage)
	flag.StringVar(&chartFile, "c", chartFile, charFileUsage)

	var CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	flag.Usage = func() {
		fmt.Fprintf(CommandLine.Output(), "Usage: %s [options]\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}

	flag.Parse()

	if *isVerbose {
		fmt.Println("Verbose...")
	}

	if *isHelp {
		flag.Usage()

		os.Exit(0)
	}

	examplers := []examples.Exampler{
		graph.TreeExamples{},
		graph.GraphExamples{
			NpmFile: chartFile,
		},
	}

	for _, e := range examplers {
		e.Examples()
	}

	fs := http.FileServer(http.Dir("web/html"))
	log.Println("running server at http://localhost:8089")
	log.Fatal(http.ListenAndServe("localhost:8089", logRequest(fs)))
}
