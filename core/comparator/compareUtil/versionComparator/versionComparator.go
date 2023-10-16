package versionComparator

import (
	"regexp"
	"strings"
)

// VersionComparator 用于比较版本号的结构体。
type VersionComparator struct{}

// Compare 比较两个版本号的方法。
//
// 如果 version1 < version2，返回负数；
// 如果 version1 = version2，返回 0；
// 如果 version1 > version2，返回正数。
//
// 版本号可以包含数字和字符，例如："1.0.0"、"1.0.2a"、"V0.0.20170102"。
func (vc *VersionComparator) Compare(version1, version2 string) int {
	if version1 == version2 {
		return 0
	}
	if version1 == "" {
		return -1
	}
	if version2 == "" {
		return 1
	}

	v1s := strings.Split(version1, ".")
	v2s := strings.Split(version2, ".")
	minLength := min(len(v1s), len(v2s))

	for i := 0; i < minLength; i++ {
		v1 := v1s[i]
		v2 := v2s[i]
		diff := len(v1) - len(v2)
		if diff == 0 {
			if v1 != v2 {
				diff = strings.Compare(v1, v2)
			} else {
				// 长度相等时，数字比较
				v1Num := extractNumber(v1)
				v2Num := extractNumber(v2)
				diff = v1Num - v2Num
			}
		}

		if diff != 0 {
			return diff
		}
	}

	return len(v1s) - len(v2s)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func extractNumber(s string) int {
	re := regexp.MustCompile(`\d+`)
	match := re.FindString(s)
	if match != "" {
		return int(match[0])
	}
	return 0
}
