package main

import (
	"flag"
	"fmt"
	"os"
	"bufio"
	"io"
	"progress/cmd/core"
	"strings"
)

func main() {
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Println("Too few args specified. Supported commands are 'times', 'unique' and 'hang'.")
		os.Exit(1)
	}

	var p core.Progress
	switch flag.Arg(0) {
	case "times":
		p = core.NewProgressTimings()
	case "unique":
		p = core.NewProgressUniques()
	case "hang":
		p = core.NewProgressHangs()
	default:
		fmt.Printf("Invalid command: '%s'. Supported commands are 'times', 'unique' and 'hang'.\n", flag.Arg(0))
		os.Exit(1)
	}

	var line string
	for i := 1; i < flag.NArg(); i++ {
		fmt.Printf("--------  %s  --------\n", flag.Arg(i))
		progressFile, err := os.Open(flag.Arg(i))
		if err != nil {
			panic(err)
		}

		reader := bufio.NewReader(progressFile)

		for {
			line, err = reader.ReadString('\n')
			line = strings.TrimSpace(line)

			p.ProcessLine(line)

			if err != nil {
				break
			}
		}

		if err != io.EOF {
			panic(err)
		}

		progressFile.Close()
	}

	p.Results()
}
