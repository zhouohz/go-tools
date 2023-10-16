package desensitized

import "strings"

// DesensitizedType is an enumeration of supported desensitization types.
type DesensitizedType int

const (
	// USERID represents the user ID.
	USERID DesensitizedType = iota
	// CHINESENAME represents the Chinese name.
	CHINESENAME
	// IDCARD represents the ID card number.
	IDCARD
	// FIXEDPHONE represents the fixed phone number.
	FIXEDPHONE
	// MOBILEPHONE represents the mobile phone number.
	MOBILEPHONE
	// ADDRESS represents the address.
	ADDRESS
	// EMAIL represents the email address.
	EMAIL
	// PASSWORD represents the password.
	PASSWORD
	// CARLICENSE represents the Chinese car license.
	CARLICENSE
	// BANKCARD represents the bank card number.
	BANKCARD
	// IPV4 represents the IPv4 address.
	IPV4
	// IPV6 represents the IPv6 address.
	IPV6
	// FIRSTMASK represents the rule to show only the first character.
	FIRSTMASK
	// CLEARTOEMPTY represents clearing to an empty string.
	CLEARTOEMPTY
	// CLEARTONULL represents clearing to nil.
	CLEARTONULL
)

// Desensitized applies desensitization using the default strategy.
func Desensitized(str string, desensitizedType DesensitizedType) string {
	if str == "" {
		return ""
	}
	newStr := str
	switch desensitizedType {
	case USERID:
		newStr = ""
	case CHINESENAME:
		newStr = chineseName(str)
	case IDCARD:
		newStr = idCardNum(str, 1, 2)
	case FIXEDPHONE:
		newStr = fixedPhone(str)
	case MOBILEPHONE:
		newStr = mobilePhone(str)
	case ADDRESS:
		newStr = address(str, 8)
	case EMAIL:
		newStr = email(str)
	case PASSWORD:
		newStr = password(str)
	case CARLICENSE:
		newStr = carLicense(str)
	case BANKCARD:
		newStr = bankCard(str)
	case IPV4:
		newStr = ipv4(str)
	case IPV6:
		newStr = ipv6(str)
	case FIRSTMASK:
		newStr = firstMask(str)
	case CLEARTOEMPTY:
		newStr = clear()
	case CLEARTONULL:
		newStr = clearToNull()
	}
	return newStr
}

// clear sets the string to an empty string.
func clear() string {
	return ""
}

// clearToNull sets the string to nil.
func clearToNull() string {
	return ""
}

// userId desensitizes the user ID.
func userId() string {
	return ""
}

// firstMask applies the first mask rule, showing only the first character.
func firstMask(str string) string {
	if str == "" {
		return ""
	}
	return string([]rune(str)[0]) + strings.Repeat("*", len(str)-1)
}

// chineseName desensitizes the Chinese name, showing only the first character.
func chineseName(fullName string) string {
	return firstMask(fullName)
}

// idCardNum desensitizes the ID card number, keeping the first `front` and last `end` digits.
func idCardNum(idCardNum string, front, end int) string {
	if idCardNum == "" {
		return ""
	}
	if front < 0 || end < 0 || (front+end) > len(idCardNum) {
		return ""
	}
	return string([]rune(idCardNum)[:front]) + strings.Repeat("*", len(idCardNum)-front-end) + string([]rune(idCardNum)[len(idCardNum)-end:])
}

// fixedPhone desensitizes the fixed phone number, showing the first 4 and last 2 digits.
func fixedPhone(num string) string {
	if num == "" {
		return ""
	}
	return string([]rune(num)[:4]) + strings.Repeat("*", len(num)-4-2) + string([]rune(num)[len(num)-2:])
}

// mobilePhone desensitizes the mobile phone number, showing the first 3 and last 4 digits.
func mobilePhone(num string) string {
	if num == "" {
		return ""
	}
	return string([]rune(num)[:3]) + strings.Repeat("*", len(num)-3-4) + string([]rune(num)[len(num)-4:])
}

// address desensitizes the address, showing up to the last `sensitiveSize` characters.
func address(address string, sensitiveSize int) string {
	if address == "" {
		return ""
	}
	length := len([]rune(address))
	if sensitiveSize >= length {
		return address
	}
	return strings.Repeat("*", length-sensitiveSize) + address[length-sensitiveSize:]
}

// email desensitizes the email address, showing the first letter of the prefix and replacing the rest with '*'.
func email(email string) string {
	if email == "" {
		return ""
	}
	atIndex := strings.Index(email, "@")
	if atIndex <= 1 {
		return email
	}
	return string([]rune(email)[0]) + strings.Repeat("*", atIndex-1) + email[atIndex:]
}

// password desensitizes the password, replacing all characters with '*'.
func password(password string) string {
	if password == "" {
		return ""
	}
	return strings.Repeat("*", len(password))
}

// carLicense desensitizes the Chinese car license, hiding the middle characters.
func carLicense(carLicense string) string {
	if carLicense == "" {
		return ""
	}
	length := len(carLicense)
	if length == 7 {
		return string([]rune(carLicense)[:3]) + strings.Repeat("*", 4) + string([]rune(carLicense)[6:])
	} else if length == 8 {
		return string([]rune(carLicense)[:3]) + strings.Repeat("*", 4) + string([]rune(carLicense)[7:])
	}
	return carLicense
}

// bankCard desensitizes the bank card number, showing the first 4 digits and masking the rest.
func bankCard(bankCardNo string) string {
	if bankCardNo == "" {
		return ""
	}
	bankCardNo = strings.ReplaceAll(bankCardNo, " ", "")
	if len(bankCardNo) < 9 {
		return bankCardNo
	}
	length := len(bankCardNo)
	endLength := length % 4
	midLength := length - 4 - endLength
	var buf strings.Builder
	buf.WriteString(bankCardNo[:4])
	for i := 0; i < midLength; i++ {
		if i%4 == 0 {
			buf.WriteRune(' ')
		}
		buf.WriteRune('*')
	}
	buf.WriteRune(' ')
	buf.WriteString(bankCardNo[length-endLength:])
	return buf.String()
}

// ipv4 desensitizes the IPv4 address, replacing the last three octets with '*'.
func ipv4(ipv4 string) string {
	parts := strings.Split(ipv4, ".")
	if len(parts) != 4 {
		return ipv4
	}
	return parts[0] + ".*.*.*"
}

// ipv6 desensitizes the IPv6 address, replacing all sections except the first with '*'.
func ipv6(ipv6 string) string {
	parts := strings.Split(ipv6, ":")
	if len(parts) < 2 {
		return ipv6
	}
	return parts[0] + ":*" + strings.Repeat(":*", len(parts)-1)
}
