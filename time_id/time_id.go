package time_id

import (
	"fmt"
	"strconv"
	"time"
)

type TimeId string

func NewByTime(date time.Time) TimeId {
	return TimeId(fmt.Sprintf("%d%02d%02d", date.Year(), date.Month(), date.Day()))
}

func (timeId TimeId) ToStr() string {
	return string(timeId)
}
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

func (timeId TimeId) YearString() string {
	return string(timeId[0:4])
}

func (timeId TimeId) MonthString() string {
	return string(timeId[4:6])
}
func (timeId TimeId) DayString() string {
	return string(timeId[6:8])
}
