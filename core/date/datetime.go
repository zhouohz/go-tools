package date

import (
	"github.com/zhouohz/go-tools/core/date/m/dateUnit"
	"github.com/zhouohz/go-tools/core/date/m/zodiac"
	"time"
)

func Time(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

// Current
//
//	@Description: 当前时间的时间戳（秒）
//	@return int64
func Current() int64 {
	return time.Now().Unix()
}

// Now
//
//	@Description: 当前时间，格式 yyyy-MM-dd HH:mm:ss
//	@return string
func Now() string {
	return FormatTimestamp(Current(), time.DateTime)
}

// ThisYear
//
//	@Description: 今年
//	@return int
func ThisYear() int {
	return Year(time.Now())
}

// ThisMonth
//
//	@Description: 当前月份
//	@return int
func ThisMonth() int {
	return Month(time.Now())
}

// ThisWeekOfYear
//
//	@Description: 当前日期所在年份的第几周
//	@return int
func ThisWeekOfYear(firstDayOfWeek time.Weekday) int {
	return WeekOfYear(time.Now(), firstDayOfWeek)
}

// ThisWeekOfMonth
//
//	@Description: 当前日期所在月份的第几周
//	@return int
func ThisWeekOfMonth() int {
	return WeekOfMonth(time.Now())
}

// ThisDayOfWeek
//
//	@Description: 当前日期是星期几
//	@return int
func ThisDayOfWeek() int {
	return DayOfWeek(time.Now())
}

// ThisHour
//
//	@Description: 当前日期的小时数部分
//	@param is24HourClock 是否24小时制
//	@return int
func ThisHour(is24HourClock bool) int {
	return Hour(time.Now(), is24HourClock)
}

// ThisMinute
//
//	@Description: 当前日期的分钟数部分
//	@return int
func ThisMinute() int {
	return Minute(time.Now())
}

// ThisSecond
//
//	@Description: 当前日期的秒数部分
//	@return int
func ThisSecond() int {
	return Second(time.Now())
}

// ThisMillisecond
//
//	@Description: 当前日期的毫秒数部分
//	@return int
func ThisMillisecond() int64 {
	return Millisecond(time.Now())
}

// DateSeconds
//
//	@Description: 获取时间（秒）
//	@return int
func DateSeconds(t time.Time) int64 {
	return t.Unix()
}

// DateMillSeconds
//
//	@Description: 获取时间（毫秒）
//	@return int64
func DateMillSeconds(t time.Time) int64 {
	return t.UnixMilli()
}

// DateNanoSeconds
//
//	@Description: 获取时间（纳秒）
//	@return int64
func DateNanoSeconds(t time.Time) int64 {
	return t.UnixNano()
}

// Year
//
//	@Description: 获得年的部分
//	@param t
//	@return int
func Year(t time.Time) int {
	return t.Year()
}

// Quarter
//
//	@Description: 获得指定日期所属季度，从1开始计数
//	@param t
//	@return int
func Quarter(t time.Time) int {
	return Month(t)/3 + 1
}

// Month
//
//	@Description: 获得月份
//	@param t
//	@return int
func Month(t time.Time) int {
	return int(t.Month())
}

// Day
//
//	@Description: 获取天
//	@param t
//	@return int
func Day(t time.Time) int {
	return t.Day()
}

// Hour
//
//	@Description: 获取小时
//	@param t
//	@return int
func Hour(t time.Time, is24HourClock bool) int {
	if is24HourClock {
		return t.Hour()
	}
	// 如果不是24小时制，将下午时段转换为上午时段
	hour := t.Hour()
	if hour >= 12 {
		hour -= 12
	}
	return hour
}

// Minute
//
//	@Description: 获取分钟
//	@param t
//	@return int
func Minute(t time.Time) int {
	return t.Hour()
}

// Second
//
//	@Description: 获取秒数
//	@param t
//	@return int
func Second(t time.Time) int {
	return t.Second()
}

// Millisecond
//
//	@Description: 获取毫秒数
//	@param t
//	@return int
func Millisecond(t time.Time) int64 {
	return t.UnixMilli()
}

// WeekOfYear
//
//	 @Description: 获得指定日期是所在年份的第几周
//		 * 此方法返回值与一周的第一天有关，比如：
//		 * 2016年1月3日为周日，如果一周的第一天为周日，那这天是第二周（返回2）
//		 * 如果一周的第一天为周一，那这天是第一周（返回1）
//		 * 跨年的那个星期得到的结果总是1
//		 *
//	 @param t
//	 @return int
func WeekOfYear(t time.Time, firstDayOfWeek time.Weekday) int {
	// 获取指定日期所在年份
	_, week := t.ISOWeek()

	// 如果一周的第一天不是星期一，需要调整周数
	if firstDayOfWeek != time.Monday {
		// 获取指定日期的星期几
		weekday := t.Weekday()

		// 计算需要调整的天数
		daysToAdjust := int(weekday - firstDayOfWeek)

		// 如果 daysToAdjust 为负数，表示需要向前调整周数
		if daysToAdjust < 0 {
			week--
		}
	}

	return week
}

// WeekOfMonth
//
//	@Description: 获得指定日期是所在月份的第几周
//	@param t
//	@return int
func WeekOfMonth(t time.Time) int {
	weekNumber := int((t.Day()-1)/7) + 1
	return weekNumber
}

// DayOfMonth
//
//	@Description: 获得指定日期是这个日期所在月份的第几天
//	@param t
//	@return int
func DayOfMonth(t time.Time) int {
	return Day(t)
}

// DayOfWeek
//
//	@Description: 获得指定日期是星期几，0表示周日，1表示周一
func DayOfWeek(t time.Time) int {
	return int(t.Weekday())
}

// IsWeekend
//
//	@Description: 是否为周末（周六或周日）
//	@param t
//	@return bool
func IsWeekend(t time.Time) bool {
	return t.Weekday() == time.Sunday || t.Weekday() == time.Saturday
}

// IsAM
//
//	@Description: 是否是上午
//	@param t
//	@return bool
func IsAM(t time.Time) bool {
	hour := t.Hour()
	return hour >= 0 && hour < 12
}

// IsPM
//
//	@Description: 是否是下午
//	@param t
//	@return bool
func IsPM(t time.Time) bool {
	hour := t.Hour()
	return hour >= 12 && hour < 24
}

// Format
//
//	@Description: 根据特定格式格式化日期
//	@return string
func Format(t time.Time, format string) string {
	return t.Format(format)
}

// FormatTimestamp
//
//	@Description: 根据特定格式格式化时间戳（秒）
//	@param timestamp
//	@param format
//	@return string
func FormatTimestamp(timestamp int64, format string) string {
	return Time(timestamp).Format(format)
}

// IsSameDay
//
//	@Description: 比较两个日期是否为同一天
//	@param t1
//	@param t2
//	@return bool
func IsSameDay(t1, t2 time.Time) bool {
	return t1.Equal(t2)
}

// IsSameWeek
//
//	@Description: 判断两个日期是否在同一周
//	@param cal1
//	@param cal2
//	@param isMon
//	@return bool
func IsSameWeek(cal1, cal2 time.Time, isMon bool) bool {
	if cal1.IsZero() || cal2.IsZero() {
		panic("日期不能为空")
	}
	first := time.Monday
	// 把所传日期设置为其当前周的第一天
	if !isMon {
		first = time.Sunday
	}
	return cal1.Year() == cal2.Year() && WeekOfYear(cal1, first) == WeekOfYear(cal2, first)
}

// IsSameMonth
//
//	@Description: 判断两个日期是否在同一个月
//	@param cal1
//	@param cal2
//	@return bool
func IsSameMonth(cal1, cal2 time.Time) bool {
	if cal1.IsZero() || cal2.IsZero() {
		panic("日期不能为空")
	}
	return cal1.Year() == cal2.Year() && cal1.Month() == cal2.Month()
}

func Offset(t time.Time, duration time.Duration, offset int) time.Time {
	return t.Add(duration * time.Duration(offset))
}

func OffsetMillisecond(t time.Time, offset int) time.Time {
	return Offset(t, time.Millisecond, offset)
}

func OffsetSecond(t time.Time, offset int) time.Time {
	return Offset(t, time.Second, offset)
}

func OffsetMinute(t time.Time, offset int) time.Time {
	return Offset(t, time.Minute, offset)
}

func OffsetHour(t time.Time, offset int) time.Time {
	return Offset(t, time.Hour, offset)
}

func OffsetDay(t time.Time, offset int) time.Time {
	return t.AddDate(0, 0, offset)
}

func OffsetWek(t time.Time, offset int) time.Time {
	return t.AddDate(0, 0, offset*7)
}

func OffsetMonth(t time.Time, offset int) time.Time {
	return t.AddDate(0, offset, 0)
}

// ------------------------------------ Offset end ----------------------------------------------

func Between(startDate, endDate time.Time, unit dateUnit.DateUnit, isAbs bool) int64 {
	if startDate.IsZero() {
		panic("startDate is null")
	}
	if endDate.IsZero() {
		panic("endDate is null")
	}

	if isAbs && startDate.After(endDate) {
		// 间隔只为正数的情况下，如果开始日期晚于结束日期，置换之
		startDate, endDate = endDate, startDate
	}

	return endDate.Sub(startDate).Milliseconds() / unit.ToMilliseconds()
}

func BetweenMs(startDate, endDate time.Time) int64 {
	return Between(startDate, endDate, dateUnit.MS, true)
}

func IsLeapYear(year int) bool {
	return ((year & 3) == 0) && ((year%100) != 0 || (year%400) == 0)
}

func GetZodiac(t time.Time) string {
	return zodiac.GetZodiac(t)
}

func GetChineseZodiac(t time.Time) string {
	return zodiac.GetChineseZodiac(t)
}

func GetChineseZodiacByYear(year int) string {
	return zodiac.GetChineseZodiacByYear(year)
}

func ConvertTimeZone(t time.Time, toZone string) (time.Time, error) {
	toLocation, err := time.LoadLocation(toZone)
	if err != nil {
		return time.Time{}, err
	}

	convertedTime := t.In(toLocation)
	return convertedTime, nil
}

func LengthOfYear(year int) int {
	// 创建指定年份的开始时间
	startTime := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)

	// 创建下一年的开始时间
	nextYear := time.Date(year+1, time.January, 1, 0, 0, 0, 0, time.UTC)

	// 计算时间差，获得总天数
	days := nextYear.Sub(startTime).Hours() / 24

	return int(days)
}

func LengthOfMonth(month time.Month, isLeapYear bool) int {
	switch month {
	case time.February: // 二月
		if isLeapYear {
			return 29
		}
		return 28
	case time.April, time.June, time.September, time.November: // 四月、六月、九月、十一月
		return 30
	default:
		return 31
	}
}

func IsLastDayOfMonth(t time.Time) bool {
	return Day(t) == LengthOfMonth(t.Month(), IsLeapYear(Year(t)))
}
