package zodiac

import "time"

// DAY_ARR 包含了每个星座更改的月份中的日期。
var DAY_ARR = []int{20, 19, 21, 20, 21, 22, 23, 23, 23, 24, 23, 22}

// ZODIACS 包含了星座的名称。
var ZODIACS = []string{
	"摩羯座", "水瓶座", "双鱼座", "白羊座", "金牛座", "双子座", "巨蟹座", "狮子座",
	"处女座", "天秤座", "天蝎座", "射手座", "摩羯座",
}

// CHINESE_ZODIACS 包含了生肖的名称。
var CHINESE_ZODIACS = []string{
	"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪",
}

// GetZodiac 通过生日计算星座。
func GetZodiac(date time.Time) string {
	return getZodiacByMonthAndDay(date.Month(), date.Day())
}

// GetZodiacByMonthAndDay 通过生日计算星座。
func getZodiacByMonthAndDay(month time.Month, day int) string {
	monthValue := int(month)
	if monthValue < 1 || monthValue > 12 {
		return ""
	}

	if day < DAY_ARR[monthValue-1] {
		return ZODIACS[monthValue-1]
	}
	return ZODIACS[monthValue]
}

// GetChineseZodiac 通过生日计算生肖，只计算1900年后出生的人。
func GetChineseZodiac(date time.Time) string {
	year := date.Year()
	if year < 1900 {
		return ""
	}
	return CHINESE_ZODIACS[(year-1900)%len(CHINESE_ZODIACS)]
}

func GetChineseZodiacByYear(year int) string {
	if year < 1900 {
		return ""
	}
	return CHINESE_ZODIACS[(year-1900)%len(CHINESE_ZODIACS)]
}
