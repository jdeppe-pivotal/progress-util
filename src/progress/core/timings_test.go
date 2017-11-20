package core_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"progress/core"
	"time"
	"sort"
)

var _ = Describe("", func() {
	It("", func() {
		p := core.NewProgressTimings()

		p.ProcessLine( "2017-11-19 15:31:34.493 +0000 Starting test TestClass testMethod")

		Expect(len(*p)).To(Equal(1))

		Expect((*p)[0].StartTime).To(BeTemporally("==", time.Date(2017, 11, 19, 15, 31, 34, int((time.Millisecond *493).Nanoseconds()), time.UTC)))

		p.ProcessLine("2017-11-19 15:31:35.493 +0000 Completed test TestClass testMethod")

		Expect(len(*p)).To(Equal(1))
		Expect((*p)[0].Duration).To(Equal(time.Second * 1))

		p.ProcessLine("2017-11-19 15:31:36.493 +0000 Starting test TestClass testMethod")
		p.ProcessLine("2017-11-19 15:31:37.493 +0000 Completed test TestClass testMethod")

		Expect(len(*p)).To(Equal(1))
		Expect((*p)[0].Duration).To(Equal(time.Second * 2))
	})

	It("Sorts", func() {
		p := core.NewProgressTimings()

		p.ProcessLine("2017-11-19 15:31:34.493 +0000 Starting test TestClass1 testMethod")
		p.ProcessLine("2017-11-19 15:31:35.493 +0000 Completed test TestClass1 testMethod")
		p.ProcessLine("2017-11-19 15:31:35.493 +0000 Starting test TestClass2 testMethod")
		p.ProcessLine("2017-11-19 15:31:37.493 +0000 Completed test TestClass2 testMethod")

		Expect(len(*p)).To(Equal(2))

		sort.Sort(p)

		Expect((*p)[0].TestClass).To(Equal("TestClass2"))
		Expect((*p)[0].Duration).To(Equal(time.Second * 2))
		Expect((*p)[0].TestCount).To(Equal(1))
		Expect((*p)[1].TestClass).To(Equal("TestClass1"))
		Expect((*p)[1].Duration).To(Equal(time.Second * 1))
		Expect((*p)[1].TestCount).To(Equal(1))
	})

	It("Accumulate", func() {
		p := core.NewProgressTimings()

		p.ProcessLine("2017-11-19 15:31:34.000 +0000 Starting test TestClass1 testMethod1")
		p.ProcessLine("2017-11-19 15:31:35.000 +0000 Completed test TestClass1 testMethod1")
		p.ProcessLine("2017-11-19 15:31:35.000 +0000 Starting test TestClass1 testMethod2")
		p.ProcessLine("2017-11-19 15:31:36.100 +0000 Completed test TestClass1 testMethod2")
		p.ProcessLine("2017-11-19 15:31:36.100 +0000 Starting test TestClass1 testMethod3")
		p.ProcessLine("2017-11-19 15:31:37.200 +0000 Completed test TestClass1 testMethod3")

		Expect(len(*p)).To(Equal(1))
		Expect((*p)[0].Duration.Seconds()).To(Equal(3.2))
		Expect((*p)[0].TestCount).To(Equal(3))
	})
})
