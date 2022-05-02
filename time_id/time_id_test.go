package time_id_test

import (
	"time"

	. "github.com/cthulhu/go-steun/time_id"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("TimeId", func() {
	It("Returns a date from timeid", func() {
		id := TimeId("20141231")
		Expect(id.ToDate()).To(Equal(time.Date(2014, time.Month(12), 31, 0, 0, 0, 0, time.UTC)))
	})

	It("Returns time ids till a given date", func() {
		start := TimeId("20141201")
		end := TimeId("20141205")
		Expect(start.Till(end)).To(Equal([]TimeId{
			TimeId("20141201"),
			TimeId("20141202"),
			TimeId("20141203"),
			TimeId("20141204"),
			TimeId("20141205"),
		}))
	})

	It("Returns a YearString from timeid", func() {
		id := TimeId("20141231")
		Expect(id.YearString()).To(Equal("2014"))
	})

	It("Returns a MonthString from timeid", func() {
		id := TimeId("20141231")
		Expect(id.MonthString()).To(Equal("12"))
	})

	It("Returns a DayString from timeid", func() {
		id := TimeId("20141231")
		Expect(id.DayString()).To(Equal("31"))
	})
	It("Returns a YearInt from timeid", func() {
		id := TimeId("20141231")
		Expect(id.YearInt()).To(Equal(2014))
	})

	It("Returns a MonthInt from timeid", func() {
		id := TimeId("20141231")
		Expect(id.MonthInt()).To(Equal(12))
	})

	It("Returns a DayInt from timeid", func() {
		id := TimeId("20141231")
		Expect(id.DayInt()).To(Equal(31))
	})
	It("Returns a String from timeid", func() {
		id := TimeId("20141231")
		Expect(id.ToStr()).To(Equal("20141231"))
	})
	It("Returns a Date from timeid", func() {
		id := TimeId("20141231")
		date, err := id.ToDate()
		Expect(err).NotTo(HaveOccurred())
		Expect(date.Year()).To(Equal(2014))
	})
	It("Returns error if date is wrong", func() {
		id := TimeId("AAAA")
		_, err := id.ToDate()
		Expect(err).To(HaveOccurred())
	})
	It("Returns error if month is wrong", func() {
		id := TimeId("2014AA12")
		_, err := id.ToDate()
		Expect(err).To(HaveOccurred())
	})
	It("Returns error if day is wrong", func() {
		id := TimeId("201412BB")
		_, err := id.ToDate()
		Expect(err).To(HaveOccurred())
	})
	Context("NewByDays", func() {
		It("returns id", func() {
			id := NewByDays(-2)
			now := time.Now().AddDate(0, 0, -2)
			Expect(id.DayInt()).To(Equal(now.Day()))
		})
	})
})
