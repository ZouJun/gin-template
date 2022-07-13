package time_helper

import "time"

//获取系统时间
func GetCurrentTimeUnix() int64 {
	return GetCurrentTime().Unix()
}

func GetCurrentTime() time.Time {
	return time.Now()
}

//昨天
func GetYesterdayTimeUnix() int64 {
	return GetYesterdayTime().Unix()
}

func GetYesterdayTime() time.Time {
	return time.Now().AddDate(0, 0, -1)
}

//月初月末
func MonthStartAndEnd(t time.Time) (time.Time, time.Time) {
	monthStart := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	monthEnd := monthStart.AddDate(0, 1, -1)

	return monthStart, monthEnd
}

//获取时间00:00:00或23:59:59
func TimeStartOrEnd(timeUnix int64, timeStart bool) time.Time {
	formatTime := time.Unix(timeUnix, 0)

	if timeStart {
		return time.Date(formatTime.Year(), formatTime.Month(), formatTime.Day(), 0, 0, 0, 0, formatTime.Location())
	} else {
		return time.Date(formatTime.Year(), formatTime.Month(), formatTime.Day(), 23, 59, 59, 0, formatTime.Location())
	}
}

func TimeStartAndEnd(timeUnix int64) (time.Time, time.Time) {
	return TimeStartOrEnd(timeUnix, true), TimeStartOrEnd(timeUnix, false)
}

//时间戳偏移
func GetShiftByTimeUnix(timeUnix int64, shift string) time.Time {
	//偏移时间量
	shiftTime, _ := time.ParseDuration(shift)
	//时间戳转time
	t := time.Unix(timeUnix, 0)
	//偏移后的时间
	return t.Add(shiftTime)
}

//日期格式化
func TimeFormat(timeUnix int64, format string) string {
	return time.Unix(timeUnix, 0).Format(format)
}

//时间偏移
func GetShiftByTime(t time.Time, shift string) time.Time {
	//偏移时间量
	shiftTime, _ := time.ParseDuration(shift)
	//偏移后的时间
	return t.Add(shiftTime)
}

//时间偏移后的时间戳
func GetShiftTimeUnit(t time.Time, shift string) int64 {
	newT := GetShiftByTime(t, shift)
	return newT.Unix()
}
