package ip

import (
	"errors"
	"net"
	"strconv"
	"strings"
)

// Convert uint to net.IP
func inet_ntoa(ipnr int64) net.IP {
	var bytes [4]byte
	bytes[0] = byte(ipnr & 0xFF)
	bytes[1] = byte((ipnr >> 8) & 0xFF)
	bytes[2] = byte((ipnr >> 16) & 0xFF)
	bytes[3] = byte((ipnr >> 24) & 0xFF)

	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}

// Convert net.IP to int64
func inet_aton(ipnr net.IP) int64 {
	bits := strings.Split(ipnr.String(), ".")

	//判断格式是否合法
	if len(bits) < 4 {
		return 0
	}

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}

// Convert string to int64
func Ipstr2Int64(ip string) (int64, error) {
	//为了防止传入的IP带有端口号，做一下处理
	index := strings.Index(ip, ":")
	newIp := ip
	if index >= 0 {
		newIp = ip[0:index]
	} else {
	}

	bits := strings.Split(newIp, ".")

	//判断格式是否合法
	if len(bits) < 4 {
		return 0, errors.New("ip format is error")
	}

	b0, err0 := strconv.Atoi(bits[0])
	if nil != err0 {
		return 0, errors.New("ip format is error")
	}

	b1, err1 := strconv.Atoi(bits[1])
	if nil != err1 {
		return 0, errors.New("ip format is error")
	}

	b2, err2 := strconv.Atoi(bits[2])
	if nil != err2 {
		return 0, errors.New("ip format is error")
	}

	b3, err3 := strconv.Atoi(bits[3])
	if nil != err3 {
		return 0, errors.New("ip format is error")
	}

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum, nil
}

//获取本地IP
func LocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()

	if nil != err {
		return "", err
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if nil != ipnet.IP.To4() {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", errors.New("can't get local IP")
}
