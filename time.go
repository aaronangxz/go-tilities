package go_tilities

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// UnixToUTC Converts current unix time to UTC time object
func UnixToUTC(timestamp int64) time.Time {
	return time.Unix(timestamp, 0).Local().UTC()
}

// ConvertTimeStampYearMonthDay Converts current unix timestamp to yyyy-mm-dd format
// time format: Mon Jan 2 15:04:05 -0700 MST 2006
func ConvertTimeStampYearMonthDay(timestamp int64, timeZone *time.Location) string {
	return fmt.Sprint(UnixToUTC(timestamp).In(timeZone).Format("2006-01-02"))
}

// ConvertTimeStampMonthDay Converts current unix timestamp to d/y format
func ConvertTimeStampMonthDay(timestamp int64, timeZone *time.Location) string {
	return fmt.Sprint(UnixToUTC(timestamp).In(timeZone).Format("2/1"))
}

// ConvertTimeStampDayOfWeekMonthDay Converts current unix timestamp to DDD dd/mm format
func ConvertTimeStampDayOfWeekMonthDay(timestamp int64, timeZone *time.Location) string {
	return fmt.Sprint(UnixToUTC(timestamp).In(timeZone).Format("Mon 02/01"))
}

// ConvertTimeStampWeekOfYear Converts current unix timestamp to the week number of year
func ConvertTimeStampWeekOfYear(timestamp int64, timeZone *time.Location) (int64, int64) {
	year, week := UnixToUTC(timestamp).In(timeZone).ISOWeek()
	return int64(year), int64(week)
}

// ConvertTimeStampTime Converts current unix timestamp to m:ss format
func ConvertTimeStampTime(timestamp int64, timeZone *time.Location) string {
	return fmt.Sprint(UnixToUTC(timestamp).In(timeZone).Format("3:04PM"))
}

func IsWeekDay(timeZone *time.Location) bool {
	day := time.Now().In(timeZone).Weekday()
	return day >= 1 && day <= 5
}

// IsStartOfWeek Checks if today is a monday
func IsStartOfWeek(timeZone *time.Location) bool {
	day := time.Now().In(timeZone).Weekday()
	return day == 1
}

// WeekStartEndDate Returns the start and end day of the current week in SGT unix time
func WeekStartEndDate(timestamp int64, timeZone *time.Location) (int64, int64) {
	date := UnixToUTC(timestamp).In(timeZone)

	startOffset := (int(time.Monday) - int(date.Weekday()) - 7) % 7
	startResult := date.Add(time.Duration(startOffset*24) * time.Hour)
	endResult := startResult.Add(time.Duration(4*24) * time.Hour)

	startYear, startMonth, startDay := startResult.Date()
	endYear, endMonth, endDay := endResult.Date()
	return time.Date(startYear, startMonth, startDay, 0, 0, 0, 0, timeZone).Unix(), time.Date(endYear, endMonth, endDay, 23, 59, 59, 59, timeZone).Unix()
}

// MonthStartEndDate Returns the start and end day of the current month in SGT unix time
func MonthStartEndDate(timestamp int64, timeZone *time.Location) (int64, int64) {
	date := UnixToUTC(timestamp).In(timeZone)
	currentYear, currentMonth, _ := date.Date()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, timeZone)
	lastOfMonth := time.Date(currentYear, currentMonth+1, 0, 23, 59, 59, 59, timeZone)
	return firstOfMonth.Unix(), lastOfMonth.Unix()
}

// YearStartEndDate Returns the start and end day of the current year in SGT unix time
func YearStartEndDate(timestamp int64, timeZone *time.Location) (int64, int64) {
	date := UnixToUTC(timestamp).In(timeZone)
	currentYear, _, _ := date.Date()
	return time.Date(currentYear, time.January, 1, 0, 0, 0, 0, timeZone).Unix(), time.Date(currentYear, time.December, 31, 23, 59, 59, 59, timeZone).Unix()
}

// GetEOD Returns 23:59:59 of today
func GetEOD(timeZone *time.Location) time.Time {
	now := time.Now().In(timeZone)
	year, month, day := time.Now().In(timeZone).Date()
	return time.Date(year, month, day, 23, 59, 59, 0, now.Location())
}

func DayStartEndDate(timestamp int64, timeZone *time.Location) (int64, int64) {
	if timeZone == nil {
		tz, _ := time.LoadLocation("UTC")
		timeZone = tz
	}
	year, month, day := UnixToUTC(timestamp).In(timeZone).Date()
	return time.Date(year, month, day, 0, 0, 0, 0, timeZone).Unix(), time.Date(year, month, day, 23, 59, 59, 59, timeZone).Unix()
}

func GetLast7DaysTSWithNow(timestamp int64, offset ...int64) (int64, int64) {
	startOfDay, _ := DayStartEndDate(timestamp, nil)
	end := startOfDay - (2 * 3600 * 24)

	if offset != nil {
		end -= offset[0]
	}

	start := end - (7 * 3600 * 24)
	return start, end
}

func GetLast28DaysTSWithNow(timestamp int64, offset ...int64) (int64, int64) {
	startOfDay, _ := DayStartEndDate(timestamp, nil)
	end := startOfDay - (2 * 3600 * 24)

	if offset != nil {
		end -= offset[0]
	}

	start := end - (28 * 3600 * 24)
	return start, end
}

func GetDayTSWithDate(date string, timeZone ...string) int64 {
	splitDate := strings.Split(date, "-")
	year, _ := strconv.Atoi(splitDate[0])
	month, _ := strconv.Atoi(splitDate[1])
	day, _ := strconv.Atoi(splitDate[2])
	m := time.Month(month)

	tz, _ := time.LoadLocation("UTC")
	if timeZone != nil {
		tz, _ = time.LoadLocation(timeZone[0])
	}
	return time.Date(year, m, day, 0, 0, 0, 0, tz).Unix() + (3600 * 24)
}
