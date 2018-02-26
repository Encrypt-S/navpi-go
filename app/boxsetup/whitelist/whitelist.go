package whitelist

// whitelist will detect and parse IP addresses
// and determine whether a given IP is contained
// within a defined range of IP addresses

// WhitelistHandler is a middleware for
// restricting requests based on whether or not they originate
// from an allowed IP range.

// type WhitelistHandler struct {
// 	allowedRanges []IpRange
// 	next          http.Handler
// }

// type IpRange struct {
// 	start net.IP
// 	end   net.IP
// }

// func newIPRange(start, end net.IP) IpRange {
// 	return IpRange{start: start, end: end}
// }

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

// func buildRanges(stringRanges []string) []IpRange {
// 	ranges := make([]IpRange, 0)
// 	for _, ip := range stringRanges {
// 		start, end := createRangesFromString(ip)
// 		ranges = append(ranges, newIPRange(start, end))
// 	}
// 	return ranges
// }

// func createRangesFromString(raw string) (net.IP, net.IP) {
// 	ips := strings.Split(raw, "-")
// 	start := net.ParseIP(ips[0])
// 	end := net.ParseIP(ips[1])
// 	return start, end
// }

// RangeMiddleware contains "Whitelisted" key that
// contains an array of allowed IP range strings
// separated by dashes
// {
//    "Allowed": [
//         "215.221.*.*-215.221.*.*",
//         "127.0.*.*-159.2.4.1",
//    ]
// }

// type RangeMiddleware struct {
// 	Allowed []string
// }

// // RangeHandler is a middleware handler that
// // restricts requests based on allowed IP range
// type RangeHandler struct {
// 	//allowedRanges []IpRange
// 	next          http.Handler
// }

//
// notes :: patterns :: possible approaches
//
//func newIPRange(start, end net.IP) IpRange {
//	return IpRange{start: start, end: end}
//}
//
//func (i IpRange) contains(ip net.IP) bool {
//	return bytes.Compare(ip, i.start) >= 0 && bytes.Compare(ip, i.end) <= 0
//}

//func (i *RangeHandler) containsIP(ip net.IP) bool {
//	for _, addrs := range i.allowedRanges {
//		if addrs.contains(ip) {
//			return true
//		}
//	}
//	return false
//}

//func buildRanges(stringRanges []string) []IpRange {
//	ranges := make([]IpRange, 0)
//	for _, ip := range stringRanges {
//		start, end := createRangesFromString(ip)
//		ranges = append(ranges, newIPRange(start, end))
//	}
//	return ranges
//}

//func createRangesFromString(raw string) (net.IP, net.IP) {
//	ips := strings.Split(raw, "-")
//	start := net.ParseIP(ips[0])
//	end := net.ParseIP(ips[1])
//	return start, end
//}
