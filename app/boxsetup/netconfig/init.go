package netconfig

// init contains the necessary functions
// required to implement the netconfig middleware interface

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/mailgun/vulcand/plugin"
	"github.com/urfave/cli"
)

const Type = "whitelist"

func GetSpec() *plugin.MiddlewareSpec {
	return &plugin.MiddlewareSpec{
		Type:      Type,
		FromOther: FromOther,
		FromCli:   FromCli,
		CliFlags:  CliFlags(),
	}
}

func New(allowed []string) (*WhitelistMiddleware, error) {
	if len(allowed) < 1 {
		return nil, fmt.Errorf("Need at least one IP range")
	}
	return &WhitelistMiddleware{Allowed: allowed}, nil
}

func (c *WhitelistMiddleware) String() string {
	return fmt.Sprintf("IP Range Handler")
}

func (c *WhitelistMiddleware) NewHandler(next http.Handler) (http.Handler, error) {
	allowedRanges := buildRanges(c.Allowed)
	return &WhitelistHandler{allowedRanges: allowedRanges, next: next}, nil
}

func FromOther(c WhitelistMiddleware) (plugin.Middleware, error) {
	return New(c.Allowed)
}

func FromCli(c *cli.Context) (plugin.Middleware, error) {
	return New(strings.Split(c.String("allowed"), ","))
}

func CliFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{"allowed, a", "", "Allowed IP ranges, in the format IP1-IP2,IP3-IP4", ""},
	}
}
