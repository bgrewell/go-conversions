package conversions

import (
	"fmt"
	"net"
)

// Inet4_aton Converts from an IP address in dotted decimal format x.x.x.x to a uint32 in in network byte order
// byte order
func Inet4_aton(ip string) (ip_int uint32) {
	ip_byte := net.ParseIP(ip).To4()
	for i := 0; i < len(ip_byte); i++ {
		ip_int |= uint32(ip_byte[i])
		if i < 3 {
			ip_int <<= 8
		}
	}
	return ip_int
}

// Inet4_haton Converts from an IP address in dotted decimal format x.x.x.x to a uint32 in network byte order
func Inet4_haton(ip string) (ip_int uint32) {
	ip_byte := net.ParseIP(ip).To4()
	for i := len(ip_byte) - 1; i >= 0; i-- {
		ip_int |= uint32(ip_byte[i])
		if i != 0 {
			ip_int <<= 8
		}
	}
	return ip_int
}

// Inet4_ntoa Converts from an IP address stored as a uint32 in network order to dotted decimal format keeping it in network byte order.
func Inet4_ntoa(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

// Inet4_ntoha Converts from an IP address stored as a uint32 in network order to dotted decimal in host order.
func Inet4_ntoha(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", byte(ip), byte(ip>>8), byte(ip>>16), byte(ip>>24))
}