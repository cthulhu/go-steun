package time_id

import (
	"fmt"
	"strconv"
	"time"
)

// TimeId - Abstraction to manipulate with days and a sequence of time ids
// For example: 20110101
type TimeId string

// NewByTime - Builds timeID form time.Time object
func NewByTime(date time.Time) TimeId {
	return TimeId(fmt.Sprintf("%d%02d%02d", date.Year(), date.Month(), date.Day()))
}

// ToStr - converts to string
func (timeId TimeId) ToStr() string {
	return string(timeId)
}

// ToDate - converts to time.Time object
func (timeId TimeId) ToDate() (time.Time, error) {
	yearInt, err := strconv.Atoi(timeId.YearString())
	if err != nil {
		return time.Time{}, err
	}
	monthInt, err := strconv.Atoi(timeId.MonthString())
	if err != nil {
		return time.Time{}, err
	}
	dayInt, err := strconv.Atoi(timeId.DayString())
	if err != nil {
		return time.Time{}, err
	}
	return time.Date(yearInt, time.Month(monthInt), dayInt, 0, 0, 0, 0, time.UTC), nil
}

// Till - returns days from TimeId to stopDate as an array of TimeIds
func (timeId TimeId) Till(stopDate TimeId) []TimeId {
	timeIds := []TimeId{}
	startDate, _ := timeId.ToDate()
	endDate, _ := stopDate.ToDate()

	for startDate.Before(endDate) {
		timeIds = append(timeIds, NewByTime(startDate))
		startDate = startDate.AddDate(0, 0, 1)
	}
	timeIds = append(timeIds, stopDate)
	return timeIds
}

// YearString - returns year from TimeId
func (timeId TimeId) YearString() string {
	return string(timeId[0:4])
}

// MonthString - returns day from TimeId
func (timeId TimeId) MonthString() string {
	return string(timeId[4:6])
}

// DayString - returns day from TimeId
func (timeId TimeId) DayString() string {
	return string(timeId[6:8])
}

// YearInt - returns year from TimeId
func (timeId TimeId) YearInt() int {
	return aToIWithDefault(timeId.YearString(), 0)
}

// MonthInt - returns month from TimeId
func (timeId TimeId) MonthInt() int {
	return aToIWithDefault(timeId.MonthString(), 0)
}

// DayInt - returns day from TimeId
func (timeId TimeId) DayInt() int {
	return aToIWithDefault(timeId.DayString(), 0)
}

func aToIWithDefault(a string, d int) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		return d
	}
	return i
}
