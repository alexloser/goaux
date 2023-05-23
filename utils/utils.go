// Utils
package utils

import "net"

func GetIPv4() []string {
	addresses, _ := net.InterfaceAddrs()
	ip_list := []string{}
	for _, address := range addresses {
		if a, ok := address.(*net.IPNet); ok && !a.IP.IsLoopback() {
			if a.IP.To4() != nil {
				ip_list = append(ip_list, a.IP.String())
			}
		}
	}
	return ip_list
}
