package net

import (
	"fmt"
	"github.com/zhouohz/go-tools/core/conv"
	"github.com/zhouohz/go-tools/core/util/number"
	"github.com/zhouohz/go-tools/core/util/str"
	"math/rand"
	"net"
	"time"
)

// Default minimum and maximum port ranges
const (
	PortRangeMin = 1024
	PortRangeMax = 65535
)

const (
	PortSegSplit = "/"
)

// IsUsableLocalPort 检测本地端口可用性
func IsUsableLocalPort(port int) bool {
	if !IsValidPort(port) {
		// 给定的IP未在指定端口范围中
		return false
	}

	// 创建一个TCP监听，检查端口是否可用
	tcpAddr := fmt.Sprintf(":%d", port)
	ln, err := net.Listen("tcp", tcpAddr)
	if err != nil {
		return false
	}
	defer ln.Close()

	// 创建一个UDP连接，检查端口是否可用
	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))
	if err != nil {
		return false
	}
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return false
	}
	defer conn.Close()

	return true
}

// GetUsableLocalPort 查找1024~65535范围内的可用端口
func GetUsableLocalPort() int {
	return GetUsableLocalPortRange(PortRangeMin, PortRangeMax)
}

// GetUsableLocalPortRange 查找指定范围内的可用端口
func GetUsableLocalPortRange(minPort, maxPort int) int {
	rand.Seed(time.Now().UnixNano())
	maxPortExclude := maxPort + 1
	for i := minPort; i < maxPortExclude; i++ {
		randomPort := minPort + rand.Intn(maxPort-minPort+1)
		if IsUsableLocalPort(randomPort) {
			return randomPort
		}
	}

	// 如果在指定范围内找不到可用端口，返回错误
	errMessage := fmt.Sprintf("Could not find an available port in the range [%d, %d] after %d attempts", minPort, maxPort, maxPort-minPort)
	panic(errMessage)
}

// PortIsRange 判断端口是否在指定范围内
func PortIsRange(port int, seg string, equal bool) bool {

	segments := str.Split(seg, PortSegSplit)

	if len(segments) != 2 {
		return false
	}
	ports := conv.ToIntSlice(segments)
	start := ports[0]
	end := ports[1]
	if IsValidPort(start) && IsValidPort(end) {
		return false
	}
	return number.NumInRange(port, start, end, equal)

}

func IsValidPort(port int) bool {
	return port >= PortRangeMin && port <= PortRangeMax
}
