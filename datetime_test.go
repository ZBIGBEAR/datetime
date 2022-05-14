package datetime

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAddSecond(t *testing.T) {
	now := time.Now()
	nowUnix := now.Second()
	add3s := AddSecond(now, 3)
	add3sUnix := add3s.Second()
	assert.Equal(t, nowUnix+3, add3sUnix)
}

func TestAddMinute(t *testing.T) {
	now := time.Now()
	nowUnix := now.Minute()
	add3m := AddMinute(now, 3)
	add3mUnix := add3m.Minute()
	assert.Equal(t, nowUnix+3, add3mUnix)
}

func TestAddHour(t *testing.T) {
	now := time.Now()
	nowUnix := now.Hour()
	add3h := AddHour(now, 3)
	add3hUnix := add3h.Hour()
	assert.Equal(t, nowUnix+3, add3hUnix)
}

func TestGetNowDate(t *testing.T) {
	t1 := GetNowDate()

	now := time.Now()
	assert.Equal(t, fmt.Sprintf("%.4d-%.2d-%.2d", now.Year(), now.Month(), now.Day()), t1)
}

func TestGetNowTime(t *testing.T) {
	t1 := GetNowTime()

	now := time.Now()
	assert.Equal(t, fmt.Sprintf("%.2d:%.2d:%.2d", now.Hour(), now.Minute(), now.Second()), t1)
}

func TestGetNowDateTime(t *testing.T) {
	t1 := GetNowDateTime()

	now := time.Now()
	assert.Equal(t, fmt.Sprintf("%.4d-%.2d-%.2d %.2d:%.2d:%.2d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second()), t1)
}

func TestStr2Time(t *testing.T) {
	strTime := "2022-05-14 02:03:04"
	t1, err := StrToTime(strTime, YYYY_MM_DD_HHMMSS)
	assert.Nil(t, err)

	t2 := TimeToStr(t1, YYYY_MM_DD_HHMMSS)
	assert.Equal(t, strTime, t2)

	t3 := TimeToStr(t1, YYYY_MM_DD)
	assert.Equal(t, "2022-05-14", t3)
}
