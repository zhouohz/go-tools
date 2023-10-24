package net

import (
	"fmt"
	"github.com/zhouohz/go-tools/core/util/str"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var localhostName string

// ToAbsoluteURL 相对URL转换为绝对URL
func ToAbsoluteURL(absoluteBasePath, relativePath string) (string, error) {
	absoluteURL, err := url.Parse(absoluteBasePath)
	if err != nil {
		return "", err
	}

	relativeURL, err := url.Parse(relativePath)
	if err != nil {
		if err != nil {
			return "", err
		}
	}

	absoluteURL = absoluteURL.ResolveReference(relativeURL)
	return absoluteURL.String(), nil
}

// GetIPByHost 通过域名得到IP
func GetIPByHost(hostName string) string {
	ipAddress, err := net.LookupIP(hostName)
	if err != nil {
		// 如果无法解析域名，返回原始域名
		return hostName
	}

	// 返回第一个解析到的IP地址
	if len(ipAddress) > 0 {
		return ipAddress[0].String()
	}

	// 如果没有找到IP地址，返回原始域名
	return hostName
}

// GetNetworkInterface 获取指定名称的网络接口
func GetNetworkInterface(name string) (*net.Interface, error) {
	interfaces, err := GetNetworkInterfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range interfaces {
		if iface.Name == name {
			return iface, nil
		}
	}

	return nil, nil
}

// GetNetworkInterfaces 获取本机所有网络接口
func GetNetworkInterfaces() ([]*net.Interface, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var result []*net.Interface
	for _, iface := range interfaces {
		result = append(result, &iface)
	}
	return result, nil
}

// GetMacAddress 获取指定地址信息中的MAC地址，使用分隔符
func GetMacAddress(inetAddress, separator string) string {
	// 获取网络接口列表
	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}

	ip := net.ParseIP(inetAddress)

	// 遍历网络接口
	for _, iface := range interfaces {
		if iface.HardwareAddr != nil && len(iface.HardwareAddr) == 6 {
			// 检查IP地址是否匹配
			addrs, err := iface.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addrs {
				ipNet, ok := addr.(*net.IPNet)
				if ok && ipNet.IP.Equal(ip) {
					mac := iface.HardwareAddr
					macStr := make([]string, 6)
					for i, b := range mac {
						macStr[i] = fmt.Sprintf("%02X", b)
					}
					return strings.Join(macStr, separator)
				}
			}
		}
	}

	return ""
}

// GetLocalHostName 获取主机名称
func GetLocalHostName() (string, error) {
	if !str.IsEmptyIfStr(localhostName) {
		return localhostName, nil
	}
	// 获取主机名
	hostName, err := os.Hostname()
	if err != nil {
		return "", err
	}
	return hostName, nil
}

// IsInRange 检查IP地址是否在CIDR规则配置范围内
func IsInRange(ip, cidr string) bool {

	// Parse the CIDR range
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		panic(fmt.Sprintf("Invalid CIDR: %s", cidr))
	}

	// Parse the IP address to be tested
	testIP := net.ParseIP(ip)
	if testIP == nil {
		panic(fmt.Sprintf("Invalid IP: %s", ip))
	}

	// Check if the IP address is within the CIDR range
	return ipNet.Contains(testIP)
}

// Ping 检测IP地址是否能 ping 通
func Ping(ip string) bool {
	return PingWithTimeout(ip, 200)
}

// PingWithTimeout 检测IP地址是否能 ping 通，带有超时设置（毫秒）
func PingWithTimeout(ip string, timeout int) bool {
	// 使用net.DialTimeout进行ping检测
	conn, err := net.DialTimeout("ip:icmp", ip, time.Duration(timeout)*time.Millisecond)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

// ParseCookies 解析Cookie信息
func ParseCookies(cookieStr string) []*http.Cookie {
	if strings.TrimSpace(cookieStr) == "" {
		return []*http.Cookie{}
	}

	header := http.Header{}
	header.Add("Cookie", cookieStr)
	request := http.Request{Header: header}

	cookies := request.Cookies()
	return cookies
}

// IsOpen 检查远程端口是否开启
func IsOpen(address *net.TCPAddr, timeout int) bool {
	conn, err := net.DialTimeout("tcp", address.String(), time.Duration(timeout)*time.Millisecond)
	if err == nil {
		conn.Close()
		return true
	}
	return false
}
