package date

import (
	"fmt"
	"testing"
	"tools/core/util/desensitized"
)

func TestDate(t *testing.T) {

	//fmt.Println(Year(time.Now()))
	//fmt.Println(Month(time.Now()))
	//fmt.Println(Day(time.Now()))
	//fmt.Println(Quarter(time.Now()))
	//fmt.Println(WeekOfMonth(time.Now()))
	//fmt.Println(IsSameWeek(time.Now(), Time(1696846174), true))
	//fmt.Println(IsSameMonth(time.Now(), Time(1696673374)))

	//fmt.Println(WeekOfYear(Time(1696760599), time.Monday))

	//fmt.Println(Between(time.Now(), time.Now().Add(-2*24*time.Hour), dateUnit.WEEK, true))
	//fmt.Println(LengthOfMonth(1, false))

	fmt.Println(desensitized.Desensitized("18090744550", desensitized.FIXEDPHONE))
}
