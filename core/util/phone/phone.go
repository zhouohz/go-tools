package phone

import (
	"regexp"
	"strings"
)

// PhoneUtil 电话号码工具类
type phoneUtil struct {
}

// NewPhoneUtil 创建一个PhoneUtil实例
func PhoneUtil() *phoneUtil {
	return &phoneUtil{}
}

// IsMobile 验证是否为中国大陆手机号码
func (pu *phoneUtil) IsMobile(value string) bool {
	return regexp.MustCompile(`^1\d{10}$`).MatchString(value)
}

// IsMobileHK 验证是否为中国香港手机号码
func (pu *phoneUtil) IsMobileHK(value string) bool {
	return regexp.MustCompile(`^5\d{7}$`).MatchString(value)
}

// IsMobileTW 验证是否为中国台湾手机号码
func (pu *phoneUtil) IsMobileTW(value string) bool {
	return regexp.MustCompile(`^09\d{8}$`).MatchString(value)
}

// IsMobileMO 验证是否为中国澳门手机号码
func (pu *phoneUtil) IsMobileMO(value string) bool {
	return regexp.MustCompile(`^6\d{7}$`).MatchString(value)
}

// IsTel 验证是否为中国大陆座机号码
func (pu *phoneUtil) IsTel(value string) bool {
	return regexp.MustCompile(`^0\d{2,1}-\d{7,8}(-\d{1,5})?$`).MatchString(value)
}

// IsTel400800 验证是否为中国大陆座机号码+400/800
func (pu *phoneUtil) IsTel400800(value string) bool {
	return regexp.MustCompile(`^0\d{2,1}-\d{7,8}(-\d{1,5})?$|^400\d{7}(\d{1,4})?$|^800\d{7}(\d{1,4})?$`).MatchString(value)
}

// IsPhone 验证是否为各种类型的电话号码
func (pu *phoneUtil) IsPhone(value string) bool {
	return pu.IsMobile(value) || pu.IsTel400800(value) || pu.IsMobileHK(value) || pu.IsMobileTW(value) || pu.IsMobileMO(value)
}

// HideBefore 隐藏手机号前7位
func (pu *phoneUtil) HideBefore(phone string) string {
	return strings.Repeat("*", 7) + phone[7:]
}

// HideBetween 隐藏手机号中间4位
func (pu *phoneUtil) HideBetween(phone string) string {
	return phone[:3] + strings.Repeat("*", 4) + phone[7:]
}

// HideAfter 隐藏手机号最后4位
func (pu *phoneUtil) HideAfter(phone string) string {
	return phone[:7] + strings.Repeat("*", 4)
}

// SubBefore 获取手机号前3位
func (pu *phoneUtil) SubBefore(phone string) string {
	return phone[:3]
}

// SubBetween 获取手机号中间4位
func (pu *phoneUtil) SubBetween(phone string) string {
	return phone[3:7]
}

// SubAfter 获取手机号后4位
func (pu *phoneUtil) SubAfter(phone string) string {
	return phone[7:]
}

// SubTelBefore 获取固话号码的区号部分
func (pu *phoneUtil) SubTelBefore(value string) string {
	re := regexp.MustCompile(`^0\d{2,1}`)
	match := re.FindString(value)
	return match
}

// SubTelAfter 获取固话号码的号码部分
func (pu *phoneUtil) SubTelAfter(value string) string {
	re := regexp.MustCompile(`-(\d{7,8}(-\d{1,5})?)?$`)
	match := re.FindStringSubmatch(value)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}
