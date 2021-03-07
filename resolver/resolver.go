package resolver

import (
	"context"
	"fmt"
	"net"
	"time"
)

func NewResolver(dnsServerIp string) *net.Resolver {
	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, addr string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}

			return d.DialContext(ctx, "tcp", fmt.Sprintf("%s:53", dnsServerIp))
		},
	}

	return resolver
}
