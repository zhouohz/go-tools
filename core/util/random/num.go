package random

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	//随机种子
	rand.Seed(time.Now().UnixNano())
}

// RandInt 生成随机整数
func RandInt(n int) int {
	if n <= 0 {
		return 0
	}
	return rand.Intn(n)
}

// RandomIntRange 生成一定范围内的随机整数
func RandomIntRange(min, max int) int {
	return rand.Intn(max-min) + min
}

// RandomFloatRange 生成一定范围内的随机浮点数
func RandomFloatRange(min, max float64) float64 {
	if min > max {
		min, max = max, min
	}
	// 到这里确保 max>=min 并且二者一定是正数
	ret := min + rand.Float64()*(max-min)
	return ret
}

// RandIntSlice 生成随机整数切片
func RandIntSlice(l int, max int) []int {
	slice := make([]int, l)
	for i := 0; i < l; i++ {
		slice[i] = RandInt(max)
	}
	return slice
}

// Interval 表示一个区间
type Interval struct {
	Start, End int
}

// RandNumberInIntervals 生成指定区间内的随机数
func RandNumberInIntervals(intervals []Interval) int {
	// 随机选择一个区间
	interval := intervals[rand.Intn(len(intervals))]

	// 生成该区间内的随机数
	randomNumber := rand.Intn(interval.End-interval.Start+1) + interval.Start

	return randomNumber
}

// RandNumberInMainInterval 生成随机数在主区间内，并剔除小区间
func RandNumberInMainInterval(mainInterval Interval, excludedIntervals []Interval) int {
	fmt.Println(mainInterval)
	fmt.Println(excludedIntervals)
	for {
		randomNumber := rand.Intn(mainInterval.End-mainInterval.Start+1) + mainInterval.Start
		valid := true
		for _, excludedInterval := range excludedIntervals {
			if randomNumber >= excludedInterval.Start && randomNumber <= excludedInterval.End {
				valid = false
				break
			}
		}
		if valid {
			return randomNumber
		}
	}
}
