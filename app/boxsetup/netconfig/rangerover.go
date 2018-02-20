package netconfig

// rangerover contains methods for parsing IP addresses
// and for determining whether a given IP is contained within
// a defined range of IP addresses.

import (
	"bytes"
	"net"
	"strings"
)

type ipRange struct {
	start net.IP
	end   net.IP
}

func newIPRange(start, end net.IP) ipRange {
	return ipRange{start: start, end: end}
}

func (i ipRange) contains(ip net.IP) bool {
	return bytes.Compare(ip, i.start) >= 0 && bytes.Compare(ip, i.end) <= 0
}

func (i *WhitelistHandler) containsIP(ip net.IP) bool {
	for _, addrs := range i.allowedRanges {
		if addrs.contains(ip) {
			return true
		}
	}
	return false
}

func buildRanges(stringRanges []string) []ipRange {
	ranges := make([]ipRange, 0)
	for _, ip := range stringRanges {
		start, end := createRangesFromString(ip)
		ranges = append(ranges, newIPRange(start, end))
	}
	return ranges
}

func createRangesFromString(raw string) (net.IP, net.IP) {
	ips := strings.Split(raw, "-")
	start := net.ParseIP(ips[0])
	end := net.ParseIP(ips[1])
	return start, end
}
