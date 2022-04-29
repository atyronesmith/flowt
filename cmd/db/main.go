package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/atyronesmith/flowt/pkg/dbparse"
	utils "github.com/atyronesmith/flowt/pkg/utils"
)

func main() {
	isVerbose := flag.Bool("verbose", false, "Print extra runtime information.")
	isHelp := flag.Bool("help", false, "Print usage information.")

	var CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	flag.Usage = func() {
		fmt.Fprintf(CommandLine.Output(), "Usage: %s [options] file_to_parse\n", filepath.Base(os.Args[0]))
		fmt.Fprintf(CommandLine.Output(), "       file_to_parse  -- Path to flow rules.\n")
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

	var filename string
	if filename = flag.Arg(0); filename == "" {
		flag.Usage()
		os.Exit(1)
	}
	
	db, _, err := dbparse.DBRead(filename)
	if err != nil {
		fmt.Printf("Error reading %s: %v",filename,err)
		os.Exit(1)
	}

	dbPretty, _ := utils.PrettyStruct(db)
	fmt.Println(dbPretty)

}
