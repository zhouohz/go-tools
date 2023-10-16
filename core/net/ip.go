package net

import (
	"net"
	"regexp"
	"strings"
)

// IsInnerIP 判定是否为内网IP
/**
 * A类 10.0.0.0-10.255.255.255
 * B类 172.16.0.0-172.31.255.255
 * C类 192.168.0.0-192.168.255.255
 */
func IsInnerIP(ipAddress string) bool {
	ip := net.ParseIP(ipAddress)
	return ip.IsPrivate()
}

// IPMatches 检测指定 IP 地址是否匹配通配符 wildcard
func IPMatches(wildcard, ipAddress string) bool {
	ipPattern := `^\d+\.\d+\.\d+\.\d+$`
	if matched, _ := regexp.MatchString(ipPattern, ipAddress); !matched {
		return false
	}

	wildcardSegments := strings.Split(wildcard, ".")
	ipSegments := strings.Split(ipAddress, ".")

	if len(wildcardSegments) != len(ipSegments) {
		return false
	}

	for i := 0; i < len(wildcardSegments); i++ {
		if wildcardSegments[i] != "*" && wildcardSegments[i] != ipSegments[i] {
			return false
		}
	}
	return true
}
