package dateModifier

import (
	"time"
)

// ModifyType 定义修改类型
type ModifyType int

const (
	TRUNCATE ModifyType = iota // 取指定日期短的起始值
	ROUND                      // 指定日期属性按照四舍五入处理
	CEILING                    // 指定日期属性按照进一法处理
)

// IGNORE_FIELDS 定义要忽略的日期字段
var IGNORE_FIELDS = []int{
	datefield.HOUR.GetValue(),                 // 与HOUR同名
	datefield.AM_PM.GetValue(),                // 此字段单独处理，不参与计算起始和结束
	datefield.DAY_OF_WEEK_IN_MONTH.GetValue(), // 不参与计算
	datefield.DAY_OF_YEAR.GetValue(),          // DAY_OF_MONTH体现
	datefield.DAY_OF_MONTH.GetValue(),         // 特殊处理
	datefield.DAY_OF_YEAR.GetValue(),          // WEEK_OF_MONTH体现
}

// Modify 修改日期
func Modify(calendar time.Time, dateField datefield.DateField, modifyType ModifyType) time.Time {
	return ModifyWithOptions(calendar, dateField, modifyType, false)
}

// ModifyWithOptions 修改日期，取起始值或者结束值
func ModifyWithOptions(calendar time.Time, dateField datefield.DateField, modifyType ModifyType, truncateMillisecond bool) time.Time {
	// AM_PM上下午特殊处理
	if dateField == datefield.AM_PM {
		isAM := isAM(calendar)
		switch modifyType {
		case TRUNCATE:
			if isAM {
				calendar = calendar.Add(-time.Hour * time.Duration(calendar.Hour()))
			} else {
				calendar = calendar.Add(-time.Hour * time.Duration(calendar.Hour())).
					Add(time.Hour * 12)
			}
		case CEILING:
			if isAM {
				calendar = calendar.Add(time.Hour*11 - time.Minute*time.Duration(calendar.Minute()) - time.Second*time.Duration(calendar.Second()))
			} else {
				calendar = calendar.Add(time.Hour*23 - time.Minute*time.Duration(calendar.Minute()) - time.Second*time.Duration(calendar.Second()))
			}
		case ROUND:
			min := cond.If[int](isAM, 0, 12)
			max := cond.If[int](isAM, 11, 23)
			href := (max-min)/2 + 1
			value := calendar.Hour()
			if value < href {
				calendar = calendar.Add(-time.Duration(calendar.Hour()) * time.Hour)
			} else {
				calendar = calendar.Add((time.Hour*12 - 1 - time.Duration(calendar.Hour())) * time.Hour)
			}
		}
		// 处理下一级别字段
		return ModifyWithOptions(calendar, dateField+1, modifyType, truncateMillisecond)
	}

	var endField int
	if truncateMillisecond {
		endField = datefield.SECOND.GetValue()
	} else {
		endField = datefield.MILLISECOND.GetValue()
	}

	// 循环处理各级字段，精确到毫秒字段
	for i := dateField.GetValue() + 1; i <= endField; i++ {
		if array.Contains[int](IGNORE_FIELDS, i) {
			// 忽略无关字段（WEEK_OF_MONTH）始终不做修改
			continue
		}

		// 在计算本周的起始和结束日时，月相关的字段忽略。
		if dateField == datefield.WEEK_OF_MONTH || dateField == datefield.WEEK_OF_YEAR {
			if i == datefield.DAY_OF_MONTH.GetValue() {
				continue
			}
		} else {
			// 其它情况忽略周相关字段计算
			if i == datefield.DAY_OF_WEEK.GetValue() {
				continue
			}
		}

		ModifyField(&calendar, i, modifyType)
	}

	if truncateMillisecond {
		calendar = calendar.Add(-time.Duration(calendar.Nanosecond()))
	}

	return calendar
}

// ModifyField 修改日期字段值
func ModifyField(calendar *time.Time, field int, modifyType ModifyType) {
	if field == datefield.HOUR.GetValue() {
		// 修正小时。HOUR为12小时制，上午的结束时间为12:00，此处改为HOUR_OF_DAY: 23:59
		field = datefield.HOUR.GetValue()
	}

	switch modifyType {
	case TRUNCATE:
		*calendar = GetBeginValue(*calendar, field)
	case CEILING:
		*calendar = GetEndValue(*calendar, field)
	case ROUND:
		min := GetBeginValue(*calendar, field)
		max := GetEndValue(*calendar, field)
		href := 0
		if field == time.DayOfWeek {
			// 星期特殊处理，假设周一是第一天，中间的为周四
			href = (min + 3) % 7
		} else {
			href = (max - min) / 2
		}
		value := GetValue(*calendar, field)
		if value < href {
			*calendar = calendar.Add(-time.Duration(value) * time.Duration(GetDuration(field)))
		} else {
			*calendar = calendar.Add(time.Duration(GetDuration(field)) - time.Duration(value)*time.Duration(GetDuration(field)))
		}
	}
}

// IsAM 检查是否是上午
func isAM(calendar time.Time) bool {
	return calendar.Hour() < 12
}
