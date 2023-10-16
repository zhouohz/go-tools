package array

import (
	"fmt"
	"math/rand"
	"tools/core/tp"
)

// INDEX_NOT_FOUND 表示未找到元素的返回值
const INDEX_NOT_FOUND = -1

func IsEmpty[T any](array []T) bool {
	return array == nil || len(array) == 0
}

func IsNotEmpty[T any](array []T) bool {
	return false == IsEmpty(array)
}

func Length[T any](array []T) int {
	if array == nil {
		return 0
	}

	return len(array)
}

// Contains 数组中是否包含元素
func Contains[T comparable](array []T, value T) bool {
	return IndexOf(array, value) > INDEX_NOT_FOUND
}

func AddAll[T comparable](arrays ...[]T) []T {
	if len(arrays) == 1 {
		return arrays[0]
	}

	// 计算总长度
	length := 0
	for _, array := range arrays {
		if len(array) > 0 {
			length += len(array)
		}
	}

	result := make([]T, length)
	length = 0
	for _, array := range arrays {
		if len(array) > 0 {
			copy(result[length:], array)
			length += len(array)
		}
	}
	return result
}

// IndexOf 在数组中查找指定元素的位置
func IndexOf[T comparable](array []T, value T) int {
	if IsNotEmpty(array) {
		for i := 0; i < len(array); i++ {
			if value == array[i] {
				return i
			}
		}
	}
	return INDEX_NOT_FOUND
}

func LastIndexOf[T comparable](array []T, value T) int {
	if IsNotEmpty(array) {
		for i := len(array) - 1; i >= 0; i-- {
			if value == array[i] {
				return i
			}
		}
	}
	return INDEX_NOT_FOUND
}

func Sub[T comparable](array []T, start, end int) []T {
	length := Length(array)
	if start < 0 {
		start += length
	}
	if end < 0 {
		end += length
	}

	if start == length {
		return []T{}
	}
	if start > end {
		start, end = end, start
	}
	if end > length {
		if start >= length {
			return []T{}
		}
		end = length
	}
	return copyOfRange(array, start, end)

}

// Remove 移除数组中的元素
func Remove[T comparable](array []T, index int) []T {
	if index < 0 || index >= len(array) {
		panic(fmt.Sprintf("Index out of range: %d", index))
	}
	return append(array[:index], array[index+1:]...)
}

// RemoveEle 移除数组中指定的元素，只会移除匹配到的第一个元素
func RemoveEle[T comparable](array []T, element T) []T {
	return Remove[T](array, IndexOf[T](array, element))
}

func Reverse[T comparable](array []T, startIndexInclusive, endIndexExclusive int) []T {
	if IsEmpty[T](array) {
		return array
	}
	i := startIndexInclusive
	j := endIndexExclusive - 1

	for j > i {
		array[i], array[j] = array[j], array[i]
		i++
		j--
	}

	return array
}

// Min 函数用于查找切片中的最小值
func Min[T tp.Number](numberArray []T) T {
	if len(numberArray) == 0 {
		panic("Number array must not be empty!")
	}
	min := numberArray[0]
	for i := 1; i < len(numberArray); i++ {
		if min > numberArray[i] {
			min = numberArray[i]
		}
	}
	return min
}

// Max 函数用于查找切片中的最大值
func Max[T tp.Number](numberArray []T) T {
	if len(numberArray) == 0 {
		panic("Number array must not be empty!")
	}
	max := numberArray[0]
	for i := 1; i < len(numberArray); i++ {
		if max < numberArray[i] {
			max = numberArray[i]
		}
	}
	return max
}

// Shuffle 打乱数组顺序，会变更原数组
//
//	@Description:
//	@param array
//	@return []T
func Shuffle[T comparable](array []T) []T {
	if IsEmpty[T](array) {
		return array
	}

	rand.Shuffle(len(array), func(i, j int) {
		array[i], array[j] = array[j], array[i]
	})
	return array
}

func Swap[T comparable](array []T, index1, index2 int) []T {
	if IsEmpty[T](array) {
		panic("array must not be empty!")
	}
	tmp := array[index1]
	array[index1] = array[index2]
	array[index2] = tmp
	return array
}

func IsSortedASC[T tp.Ordered](array []T) bool {
	if IsEmpty[T](array) {
		return false
	}

	for i := 0; i < len(array)-1; i++ {
		if array[i] > array[i+1] {
			return false
		}
	}

	return true
}

func IsSortedDESC[T tp.Ordered](array []T) bool {
	if IsEmpty[T](array) {
		return false
	}

	for i := 0; i < len(array)-1; i++ {
		if array[i] < array[i+1] {
			return false
		}
	}

	return true
}

// copyOfRange 复制数组的一部分
func copyOfRange[T any](original []T, from, to int) []T {
	newLength := to - from
	if newLength < 0 {
		panic(fmt.Sprintf("%v > %v", from, to))
	}
	c := make([]T, newLength)
	copyLen := len(original) - from
	if copyLen > newLength {
		copyLen = newLength
	}
	c = append(c, original[from:from+copyLen]...)
	return c
}
