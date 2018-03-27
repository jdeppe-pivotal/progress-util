package main

import (
	"flag"
	"os"
	"bufio"
	"io"
	"strings"
	"progress/core"
)

func main() {
	flag.Parse()

	var p core.Progress
	switch flag.Arg(0) {
	case "times":
		p = core.NewProgressTimings()
	case "unique":
		p = core.NewProgressUniques()
	case "hang":
		p = core.NewProgressHangs()
	}

	var line string
	for i := 1; i < flag.NArg(); i++ {
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
