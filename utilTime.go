package serverutils

import "time"

const TIME_FORMAT_CLIENT string = "2006-01-02 03:04:05"
const TIME_FORMAT_LOG_FILE string = "2006-01-02_030405"
const TIME_FORMAT_DATE string = "2006-01-02"

type timeUtil struct {
}

//Return the time at 00:00:00 of the same day as givenTime
func (*timeUtil) GetTimeBeginOfDay(givenTime time.Time) time.Time {
	year, month, day := givenTime.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, givenTime.Location())
}

//Return the time at 00:00:00 of the day before givenTime
func (*timeUtil) GetTimeBeginOfYesterday(givenTime time.Time) time.Time {
	year, month, day := givenTime.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, givenTime.Location()).Add(time.Hour * 24 * -1)
}

func (u *timeUtil) GetSecondsSinceBeginOfDay(t time.Time) int64 {
	bod := u.GetTimeBeginOfDay(t)
	result := t.Sub(bod).Seconds()
	return int64(result)
}

func (*timeUtil) GetTimeBeginOfNextDay(t time.Time) time.Time {
	year, month, day := t.Date()
	bod := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	return bod.Add(time.Duration(24) * time.Hour)
}

func (*timeUtil) IsDifferentDay(t1 time.Time, t2 time.Time) bool {
	day1 := time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, t1.Location())
	day2 := time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, t2.Location())
	return !day2.Equal(day1)
}

func (*timeUtil) CalculateDayPassed(from time.Time, to time.Time) int {
	dayFrom := time.Date(from.Year(), from.Month(), from.Day(), 0, 0, 0, 0, from.Location())
	dayTo := time.Date(to.Year(), to.Month(), to.Day(), 0, 0, 0, 0, to.Location())
	return int(dayTo.Sub(dayFrom).Hours() / 24)
}

func (*timeUtil) FormatTimeLogTime(from time.Time) string {
	return from.Format(TIME_FORMAT_LOG_FILE)
}

/*
Return formatted time for log file, only has digits and _
*/
func (*timeUtil) FormatTimeLogFile(from time.Time) string {
	return from.Format(TIME_FORMAT_LOG_FILE)
}

/*
Return formatted date as YYYY-MM-DD
*/
func (*timeUtil) FormatTimeDate(from time.Time) string {
	return from.Format(TIME_FORMAT_DATE)
}
