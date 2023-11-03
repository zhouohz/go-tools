package number

import (
	"fmt"
	"github.com/zhouohz/go-tools/core/util/str"
	"math"
	"strconv"
	"strings"
)

// DivWithInt 两个整数相除,保留小数
func DivWithInt(n1, n2, p int) float64 {
	r, _ := strconv.ParseFloat(fmt.Sprintf(DecimalPointFormat(p), float64(n1)/float64(n2)), 64)
	return r
}

// DivWithFloat 两个float相除,保留小数
func DivWithFloat(n1, n2, float, p int) float64 {
	r, _ := strconv.ParseFloat(fmt.Sprintf(DecimalPointFormat(p), float64(n1)/float64(n2)), 64)
	return r
}

// DecimalPointFormat 获取小数点的格式化
func DecimalPointFormat(p int) string {
	if p == 0 {
		p = 2
	}
	return str.StrConcat("%.", strconv.Itoa(p), "f")
}

// NumFillZero 数字转字符串,位数不够的前面补0
func NumFillZero(n, l int) string {
	numStr := strconv.Itoa(n)
	nl := len(numStr)
	if nl >= l {
		return numStr
	}
	sb := strings.Builder{}
	for i := 0; i < (l - nl); i++ {
		sb.WriteString("0")
	}
	sb.WriteString(numStr)
	return sb.String()
}

// NumMulti 是否一个数字是否是另一个数字的整数倍
func NumMulti(n1, n2 int) bool {
	return (n1 % n2) == 0
}

// NumInRange 是否一个数字在两个数字之间
func NumInRange(in int, rangeStart, rangeEnd int, equal bool) bool {
	if equal {
		return in >= rangeStart && in <= rangeEnd
	}

	return in > rangeStart && in < rangeEnd
}

// CalculateHypotenuse 计算直角三角形的斜边长度并向上取整
func CalculateHypotenuse(a, b int) int {
	// 使用整数运算计算斜边长度的平方
	cSquared := a*a + b*b

	// 向上取整
	hypotenuse := int(math.Ceil(math.Sqrt(float64(cSquared))))

	return hypotenuse
}
