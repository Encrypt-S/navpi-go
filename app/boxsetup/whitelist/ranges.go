package whitelist

// import (
// 	"net"
// 	"strings"
// )

// TODO: ranges.go contains methods for parsing IP addresses
// and for determining whether a given IP is contained within
// a defined range of IP addresses.

// import (
// 	"bytes"
// 	"net"
// 	"strings"
// )

//type IpRange struct {
//	start net.IP
//	end   net.IP
//}
//
//func newIPRange(start, end net.IP) IpRange {
//	return IpRange{start: start, end: end}
//}

// func (i IpRange) contains(ip net.IP) bool {
// 	return bytes.Compare(ip, i.start) >= 0 && bytes.Compare(ip, i.end) <= 0
// }

// func (i *WhitelistHandler) containsIP(ip net.IP) bool {
// 	for _, addrs := range i.allowedRanges {
// 		if addrs.contains(ip) {
// 			return true
// 		}
// 	}
// 	return false
// }
//
// func BuildRanges(stringRanges []string) []IpRange {
// 	ranges := make([]IpRange, 0)
// 	for _, ip := range stringRanges {
// 		start, end := CreateRangesFromString(ip)
// 		ranges = append(ranges, newIPRange(start, end))
// 	}
// 	return ranges
// }
//
// func CreateRangesFromString(raw string) (net.IP, net.IP) {
// 	ips := strings.Split(raw, "-")
// 	start := net.ParseIP(ips[0])
// 	end := net.ParseIP(ips[1])
// 	return start, end
// }
