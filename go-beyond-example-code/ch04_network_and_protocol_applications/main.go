package main

import "fmt"

// IP address manipulation
func ipToInt(ip string) uint32 {
	// Simplified - assumes valid IP format
	var a, b, c, d uint32
	fmt.Sscanf(ip, "%d.%d.%d.%d", &a, &b, &c, &d)
	return a<<24 | b<<16 | c<<8 | d
}

func intToIP(ip uint32) string {
	a := uint8(ip >> 24)
	b := uint8(ip >> 16)
	c := uint8(ip >> 8)
	d := uint8(ip)
	return fmt.Sprintf("%d.%d.%d.%d", a, b, c, d)
}

// Network mask operations
func isInSubnet(ip, network, mask uint32) bool {
	return (ip & mask) == (network & mask)
}

func main() {
	ipStr := "192.168.1.1"
	ipInt := ipToInt(ipStr)
	fmt.Printf("IP '%s' as integer: %d (0x%08X)\n", ipStr, ipInt, ipInt)

	ipBack := intToIP(ipInt)
	fmt.Printf("Integer %d as IP: %s\n", ipInt, ipBack)

	ip := ipToInt("192.168.1.100")
	network := ipToInt("192.168.1.0")
	mask := ipToInt("255.255.255.0") // /24 subnet

	inSubnet := isInSubnet(ip, network, mask)
	fmt.Printf("IP %s in subnet %s/%d: %t\n",
		intToIP(ip), intToIP(network), 24, inSubnet)
}
