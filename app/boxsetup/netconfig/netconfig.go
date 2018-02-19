package netconfig

// netconfig will detect and parse IP addresses
// and determine whether a given IP is contained
// within a defined range of IP addresses

import (
	"log"
	"net"
	"net/http"
	"io"
)

//type ipRange struct {
//	start net.IP
//	end   net.IP
//}

// RangeMiddleware contains "Whitelisted" key that
// contains an array of allowed IP range strings
// separated by dashes
// {
//    "Allowed": [
//         "215.221.*.*-215.221.*.*",
//         "127.0.*.*-159.2.4.1",
//    ]
// }

type RangeMiddleware struct {
	Allowed []string
}

// RangeHandler is a middleware handler that
// restricts requests based on allowed IP range
type RangeHandler struct {
	allowedRanges []ipRange
	next          http.Handler
}

func HttpScan(w http.ResponseWriter, r *http.Request) {

	host, port, err := net.SplitHostPort(r.RemoteAddr)

	log.Println("host", host)
	log.Println("port", port)
	log.Println("err", err)

	if err != nil || host == "" {
		// w.WriteHeader(http.StatusInternalServerError)
		requestIP := net.ParseIP(host)
		log.Println(requestIP)
	}

	// if we are not in localhost parse the IP
	if host != "::1" {
		requestIP := net.ParseIP(host)
		if !i.containsIP(requestIP) {
			w.WriteHeader(http.StatusForbidden)
			io.WriteString(w, "Forbidden")
			return
		}
	}

}

//
// notes :: patterns :: possible approaches
//
//func newIPRange(start, end net.IP) ipRange {
//	return ipRange{start: start, end: end}
//}
//
//func (i ipRange) contains(ip net.IP) bool {
//	return bytes.Compare(ip, i.start) >= 0 && bytes.Compare(ip, i.end) <= 0
//}
//
//func (i *RangeHandler) containsIP(ip net.IP) bool {
//	for _, addrs := range i.allowedRanges {
//		if addrs.contains(ip) {
//			return true
//		}
//	}
//	return false
//}
//
//func buildRanges(stringRanges []string) []ipRange {
//	ranges := make([]ipRange, 0)
//	for _, ip := range stringRanges {
//		start, end := createRangesFromString(ip)
//		ranges = append(ranges, newIPRange(start, end))
//	}
//	return ranges
//}
//
//func createRangesFromString(raw string) (net.IP, net.IP) {
//	ips := strings.Split(raw, "-")
//	start := net.ParseIP(ips[0])
//	end := net.ParseIP(ips[1])
//	return start, end
//}
