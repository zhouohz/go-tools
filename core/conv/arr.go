package conv

import "time"

// ToSlice 将接口数据类型转换为[]interface{}
func ToSlice(i interface{}) []interface{} {
	v, _ := ToSliceE(i)
	return v
}

// ToBoolSlice 将接口数据类型转换为[]bool
func ToBoolSlice(i interface{}) []bool {
	v, _ := ToBoolSliceE(i)
	return v
}

// ToStrSlice 将接口数据类型转换为[]string
func ToStrSlice(i interface{}) []string {
	v, _ := ToStrSliceE(i)
	return v
}

// ToIntSlice 将接口数据类型转换为[]int
func ToIntSlice(i interface{}) []int {
	v, _ := ToIntSliceE(i)
	return v
}

// ToDurationSlice 将接口数据类型转换为[]time.Duration
func ToDurationSlice(i interface{}) []time.Duration {
	v, _ := ToDurationSliceE(i)
	return v
}
