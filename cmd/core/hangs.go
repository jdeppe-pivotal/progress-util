package core

import (
	"strings"
	"fmt"
)

type ProgressHangs struct {
	tests map[string]string
	startTime *string
	endTime *string
}

func NewProgressHangs() *ProgressHangs {
	ph := &ProgressHangs{}
	ph.tests = make(map[string]string)
	return ph
}

func (this *ProgressHangs) ProcessLine(line string) {
	parts := strings.Split(line, " ")
	if len(parts) < 7 {
		return
	}

	timeStamp := fmt.Sprintf("%s %s %s", parts[0], parts[1], parts[2])
	this.endTime = &timeStamp

	testName := fmt.Sprintf("%s %s", parts[5], parts[6])
	state := parts[3]

	if state == "Starting" {
		this.tests[testName] = timeStamp

		if this.startTime == nil {
			this.startTime = &timeStamp
		}
	} else if state == "Completed" {
		delete(this.tests, testName)
	}
}

func (this *ProgressHangs) Results() {
	fmt.Printf("Started @ %s\n", *this.startTime)
	for k, v := range this.tests {
		fmt.Printf("%s  %s\n", v, k)
	}
	fmt.Printf("Ended @ %s\n", *this.endTime)

}

