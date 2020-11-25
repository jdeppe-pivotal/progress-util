package core

import (
	"strings"
	"fmt"
)

type ProgressUniques struct {
	tests []string
}

func NewProgressUniques() *ProgressUniques {
	t := make([]string, 0)
	return &ProgressUniques{tests: t}
}

func (this *ProgressUniques) ProcessLine(line string) {
	parts := strings.Split(line, " ")
	if len(parts) < 7 {
		return
	}

	testName := parts[5]
	state := parts[3]

	if state == "Starting" && ! entryExists(this.tests, testName) {
		this.tests = append(this.tests, testName)
	}
}

func (this *ProgressUniques) Results() {
	for _, v := range this.tests {
		fmt.Println(v)
	}
}

func entryExists(list []string, entry string) bool {
	for _, v := range list {
		if v == entry {
			return true
		}
	}
	return false
}

