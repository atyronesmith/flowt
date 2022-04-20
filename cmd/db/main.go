package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/atyronesmith/flowt/pkg/types"
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

	var in io.Reader

	var filename string

	if filename = flag.Arg(0); filename != "" {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println("error opening file: err:", err)
			os.Exit(1)
		}
		defer f.Close()

		in = f
	} else {
		flag.Usage()
		os.Exit(1)
	}
	stats, _ := os.Stat(filename)

	scanner := bufio.NewScanner(in)

	scanner.Buffer(make([]byte, 0), int(stats.Size()))
	scanner.Split(bufio.ScanLines)

	// OVSDB JSON 867056 d03a2cb121de101e6f9a1906751f54f05cd80606
	// {
	// ...
	// }

	jsonStart := regexp.MustCompile(`OVSDB\s+JSON\s+\d+\s+[0-9a-f]+`)

	lineNo := 1

	inJson := false

	var jsonObject strings.Builder

	data := types.LogicalSwitchPortTable{}

	for scanner.Scan() {
		line := scanner.Text()

		if jsonStart.MatchString(line) {
			if inJson {
				if err := json.Unmarshal([]byte(jsonObject.String()), &data); err != nil {
					fmt.Printf("Error while unmarshalling: %v", err)
				}

				if len(data.LogicalSwitchPort) > 0 {
					break
				}
				//process
				// break if true
			}
			inJson = true
			jsonObject.Reset()
		} else {
			jsonObject.WriteString(line)
		}
		lineNo += 1
	}

	fmt.Printf("%+v\n",data)
	// for _, lsp := range data.LogicalSwitchPort {
	// 	//		for _, ex := range lsp.ExternalIds {
	// 	fmt.Printf("%+v\n", lsp)
	// 	//		}
	// }

	//	file, _ := ioutil.ReadAll(in)

}
