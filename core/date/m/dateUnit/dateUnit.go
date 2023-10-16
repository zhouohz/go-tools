package dateUnit

// DateUnit 代表一个自定义的日期单位类型
type DateUnit int64

// 定义不同日期单位的常量（以毫秒为基准）
const (
	MS     DateUnit = 1
	SECOND DateUnit = 1000
	MINUTE          = 60 * SECOND
	HOUR            = 60 * MINUTE
	DAY             = 24 * HOUR
	WEEK            = 7 * DAY
)

// ToMilliseconds 将DateUnit转换为毫秒
func (du DateUnit) ToMilliseconds() int64 {
	return int64(du)
}

// ToDateUnit 将毫秒转换为DateUnit
func ToDateUnit(milliseconds int64) DateUnit {
	return DateUnit(milliseconds)
}
