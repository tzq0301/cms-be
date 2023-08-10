package ip

import "net"

// Get 获取 IPv4 和 IPv6
//
// Note: 不一定能获取到公网，没做测试，待完善
func Get() (ipv4, ipv6 string, err error) {
	ifAddrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, address := range ifAddrs {
		if ipnet, ok := address.(*net.IPNet); ok && ipnet.IP.To4() != nil {
			ipv4 = ipnet.IP.String()
		}

		if ipnet, ok := address.(*net.IPNet); ok && ipnet.IP.To4() == nil {
			ipv6 = ipnet.IP.String()
		}
	}

	return
}
