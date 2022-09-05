package serverutils

import "time"

func GetTimeBeginOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func GetTimeBeginOfYesterday(t time.Time) time.Time {
	year, month, day := t.Date()

	return time.Date(year, month, day, 0, 0, 0, 0, t.Location()).Add(time.Hour * 24 * -1)
}

func GetSecondsSinceBeginOfDay(t time.Time) int64 {
	bod := GetTimeBeginOfDay(t)
	result := t.Sub(bod).Seconds()
	return int64(result)
}

func GetTimeBeginOfNextDay(t time.Time) time.Time {
	year, month, day := t.Date()
	bod := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	return bod.Add(time.Duration(24) * time.Hour)
}

func IsDifferentDay(t1 time.Time, t2 time.Time) bool {
	day1 := time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, t1.Location())
	day2 := time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, t2.Location())
	return !day2.Equal(day1)
}

func CalculateDayPassed(from time.Time, to time.Time) int {
	dayFrom := time.Date(from.Year(), from.Month(), from.Day(), 0, 0, 0, 0, from.Location())
	dayTo := time.Date(to.Year(), to.Month(), to.Day(), 0, 0, 0, 0, to.Location())
	return int(dayTo.Sub(dayFrom).Hours() / 24)
}
