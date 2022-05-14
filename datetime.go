package datetime

import (
	"fmt"
	"time"
)

type DateTimeFormatType string

const (
	YYYY_MM_DD_HHMMSS DateTimeFormatType = "yyyy-mm-dd hh:mm:ss"
	YYYY_MM_DD_HHMM   DateTimeFormatType = "yyyy-mm-dd hh:mm"
	YYYY_MM_DD_HH     DateTimeFormatType = "yyyy-mm-dd hh"
	YYYY_MM_DD        DateTimeFormatType = "yyyy-mm-dd"
	YYYY_MM           DateTimeFormatType = "yyyy-mm"
	MM_DD             DateTimeFormatType = "mm-dd"
	DD_MM_YY_HHMMSS   DateTimeFormatType = "dd-mm-yy hh:mm:ss"

	YYYYMMDD_HHMMSS DateTimeFormatType = "yyyy/mm/dd hh:mm:ss"
	YYYYMMDD_HHMM   DateTimeFormatType = "yyyy/mm/dd hh:mm"
	YYYYMMDD_HH     DateTimeFormatType = "yyyy/mm/dd hh"
	YYYYMMDD        DateTimeFormatType = "yyyy/mm/dd"
	YYYYMM          DateTimeFormatType = "yyyy/mm"
	MMDD            DateTimeFormatType = "mm/dd"
	DDMMYY_HHMMSS   DateTimeFormatType = "dd/mm/yy hh:mm:ss"
	YYYY            DateTimeFormatType = "yyyy"
	MM              DateTimeFormatType = "mm"
	HHMMSS          DateTimeFormatType = "hh:mm:ss"
	MMSS            DateTimeFormatType = "mm:ss"
)

var timeFormat map[DateTimeFormatType]string

func init() {
	timeFormat = map[DateTimeFormatType]string{
		YYYY_MM_DD_HHMMSS: "2006-01-02 15:04:05",
		YYYY_MM_DD_HHMM:   "2006-01-02 15:04",
		YYYY_MM_DD_HH:     "2006-01-02 15:04",
		YYYY_MM_DD:        "2006-01-02",
		YYYY_MM:           "2006-01",
		MM_DD:             "01-02",
		DD_MM_YY_HHMMSS:   "02-01-06 15:04:05",
		YYYYMMDD_HHMMSS:   "2006/01/02 15:04:05",
		YYYYMMDD_HHMM:     "2006/01/02 15:04",
		YYYYMMDD_HH:       "2006/01/02 15",
		YYYYMMDD:          "2006/01/02",
		YYYYMM:            "2006/01",
		MMDD:              "01/02",
		DDMMYY_HHMMSS:     "02/01/06 15:04:05",
		YYYY:              "2006",
		MM:                "01",
		HHMMSS:            "15:04:05",
		MMSS:              "04:05",
	}
}

// AddSecond add or sub second to the time
func AddSecond(t time.Time, second int64) time.Time {
	return t.Add(time.Duration(second) * time.Second)
}

// AddMinute add or sub minute to the time
func AddMinute(t time.Time, minute int64) time.Time {
	return t.Add(time.Duration(minute) * time.Minute)
}

// AddHour add or sub hour to the time
func AddHour(t time.Time, hour int64) time.Time {
	return t.Add(time.Duration(hour) * time.Hour)
}

// AddDay add or sub day to the time
func AddDay(t time.Time, day int64) time.Time {
	return t.AddDate(0, 0, int(day))
}

// AddMonth add or sub month to the time
func AddMonth(t time.Time, month int64) time.Time {
	return t.AddDate(0, int(month), 0)
}

// AddYear add or sub year to the time
func AddYear(t time.Time, year int64) time.Time {
	return t.AddDate(int(year), 0, 0)
}

// GetNowDate return format yyyy-mm-dd of current date
func GetNowDate() string {
	return time.Now().Format(string(timeFormat[YYYY_MM_DD]))
}

// GetNowTime return format hh-mm-ss of current time
func GetNowTime() string {
	return time.Now().Format(string(timeFormat[HHMMSS]))
}

// GetNowDateTime return format yyyy-mm-dd hh-mm-ss of current datetime
func GetNowDateTime() string {
	return time.Now().Format(string(timeFormat[YYYY_MM_DD_HHMMSS]))
}

// GetNowUnix return unix of current datetime
func GetNowUnix() int64 {
	return time.Now().Unix()
}

// TimeToStr convert time to string
func TimeToStr(t time.Time, format DateTimeFormatType) string {
	return t.Format(timeFormat[format])
}

// StrToTime convert string to time
func StrToTime(str string, format DateTimeFormatType) (time.Time, error) {
	v, ok := timeFormat[format]
	if !ok {
		return time.Time{}, fmt.Errorf("format %s not found", format)
	}

	return time.Parse(v, str)
}

// BeginOfMinute return beginning minute time of day
func BeginOfMinute(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), t.Minute(), 0, 0, t.Location())
}

// EndOfMinute return end minute time of day
func EndOfMinute(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), t.Minute(), 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginOfHour return beginning hour time of day
func BeginOfHour(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 0, 0, 0, t.Location())
}

// EndOfHour return end hour time of day
func EndOfHour(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginOfDay return beginning hour time of day
func BeginOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

// EndOfDay return end time of day
func EndOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginOfWeek return beginning week, week begin from Sunday
func BeginOfWeek(t time.Time) time.Time {
	y, m, d := t.AddDate(0, 0, 0-int(BeginOfDay(t).Weekday())).Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

// EndOfWeek return end week time, week end with Saturday
func EndOfWeek(t time.Time) time.Time {
	y, m, d := BeginOfWeek(t).AddDate(0, 0, 7).Add(-time.Nanosecond).Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), t.Location())
}

// BeginOfMonth return beginning of month
func BeginOfMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth return end of month
func EndOfMonth(t time.Time) time.Time {
	return BeginOfMonth(t).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// BeginOfYear return beginning of year
func BeginOfYear(t time.Time) time.Time {
	y, _, _ := t.Date()
	return time.Date(y, time.January, 1, 0, 0, 0, 0, t.Location())
}

// EndOfYear return end of year
func EndOfYear(t time.Time) time.Time {
	return BeginOfYear(t).AddDate(1, 0, 0).Add(-time.Nanosecond)
}
