package netconfig

import (
	"io"
	"net"
	"net/http"
)

// WhitelistMiddleware holds the configuration parameters -
// it expects a "Allowed" key that is an array of strings
// separated by dashes, e.g.:
// {
//    "Allowed": [
//         "215.221.102.100-215.221.103.0",
//         "168.2.2.1-159.2.4.1",
//    ]
// }
type WhitelistMiddleware struct {
	Allowed []string
}

// WhitelistHandler is a vulcand-compatible middleware for
// restricting requests based on whether or not they originate
// from an allowed IP range.
type WhitelistHandler struct {
	allowedRanges []ipRange
	next          http.Handler
}

// ServeHTTP as implemented by WhitelistHandlers checks the request's remote address
// to see whether it is either localhost or an IP address within the configured
// whitelist ranges.
func (i *WhitelistHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host, _, err := net.SplitHostPort(r.RemoteAddr)

	if err != nil || host == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if host != "::1" {
		requestIP := net.ParseIP(host)
		if !i.containsIP(requestIP) {
			w.WriteHeader(http.StatusForbidden)
			io.WriteString(w, "Forbidden")
			return
		}
	}

	i.next.ServeHTTP(w, r)
}
