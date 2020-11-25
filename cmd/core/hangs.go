package core

import (
	"strings"
	"fmt"
)

type ProgressHangs struct {
	tests map[string][]string
	startTime *string
	endTime *string
}

func NewProgressHangs() *ProgressHangs {
	ph := &ProgressHangs{}
	ph.tests = make(map[string][]string)
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
		var val []string
		var ok bool
		if val, ok = this.tests[testName]; !ok {
			val = make([]string, 0)
		}

		val = append(val, timeStamp)
		this.tests[testName] = val

		if this.startTime == nil {
			this.startTime = &timeStamp
		}
	} else if state == "Completed" {
		val := this.tests[testName]
		lenVal := len(val)

		if lenVal == 1 {
			delete(this.tests, testName)
		} else {
			val = val[0:lenVal-1]
			this.tests[testName] = val
		}
	}
}

func (this *ProgressHangs) Results() {
	fmt.Printf("Started @ %s\n", *this.startTime)
	for k, v := range this.tests {
		for _, t := range v {
			fmt.Printf("%s  %s\n", t, k)
		}
	}
	fmt.Printf("Ended @ %s\n", *this.endTime)

}

