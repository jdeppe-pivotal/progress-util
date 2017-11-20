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

	progressFile, err := os.Open(flag.Arg(1))
	if err != nil {
		panic(err)
	}

	defer progressFile.Close()

	reader := bufio.NewReader(progressFile)

	var line string
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

	p.Results()
}

