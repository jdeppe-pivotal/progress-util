package core

import (
	"strings"
	"fmt"
	"time"
	"log"
	"sort"
)

type oneTest struct {
	TestClass string
	StartTime time.Time
	EndTime   time.Time
	Duration  time.Duration
	TestCount int
}

type ProgressTimings []*oneTest

func NewProgressTimings() *ProgressTimings {
	return &ProgressTimings{}
}

// 2017-11-19 15:31:34.493 +0000 Starting test org....
func (this *ProgressTimings) ProcessLine(line string) {
	parts := strings.Split(line, " ")
	if len(parts) < 7 {
		return
	}

	timeStampRaw := fmt.Sprintf("%s %s %s", parts[0], parts[1], parts[2])

	timeStamp, err := time.Parse("2006-01-02 15:04:05.000 -0700", timeStampRaw)
	if err != nil {
		log.Panicf("Unable to process timestamp %s", timeStampRaw)
	}

	testClass := parts[5]
	state := parts[3]

	entry, found := this.get(testClass)
	if state == "Starting" {
		if found {
			entry.StartTime = timeStamp
			entry.TestCount += 1
		} else {
			*this = append(*this, &oneTest{
				TestClass: testClass,
				StartTime: timeStamp,
				TestCount: 1,
			})
		}
	} else if state == "Completed" {
		if !found {
			log.Panicf("Unable to find entry for %s", testClass)
		}
		entry.EndTime = timeStamp
		entry.Duration += entry.EndTime.Sub(entry.StartTime)
	}
}

func (this *ProgressTimings) Results() {
	// First get the total so we can calculate %/test
	totalDuration := 0.0
	for _, v := range *this {
		totalDuration += v.Duration.Seconds()
	}

	sort.Sort(this)

	totalTestCount := 0
	for _, v := range *this {
		fmt.Printf("%9.3f  %4d  %3.2f  %s\n", v.Duration.Seconds(), v.TestCount, (v.Duration.Seconds()/totalDuration)*100, v.TestClass)
		totalTestCount += v.TestCount
	}

	fmt.Printf("%9.3f  %4d\n", totalDuration, totalTestCount)
}

func (this *ProgressTimings) Len() int {
	return len(*this)
}

func (this *ProgressTimings) Swap(i, j int) {
	(*this)[i], (*this)[j] = (*this)[j], (*this)[i]
}

func (this *ProgressTimings) Less(i, j int) bool {
	return (*this)[i].Duration.Nanoseconds() > (*this)[j].Duration.Nanoseconds()
}

func (this *ProgressTimings) get(key string) (*oneTest, bool) {
	for _, v := range *this {
		if v.TestClass == key {
			return v, true
		}
	}

	return nil, false
}
