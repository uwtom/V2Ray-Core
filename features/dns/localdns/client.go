package localdns

import (
	"context"
	"net"

	"v2ray.com/core/features/dns"
)

// Client is an implementation of dns.Client, which queries localhost for DNS.
type Client struct {
	resolver net.Resolver
}

// Type implements common.HasType.
func (*Client) Type() interface{} {
	return dns.ClientType()
}

// Start implements common.Runnable.
func (*Client) Start() error { return nil }

// Close implements common.Closable.
func (*Client) Close() error { return nil }

// LookupIP implements Client.
func (c *Client) LookupIP(host string) ([]net.IP, error) {
	ipAddr, err := c.resolver.LookupIPAddr(context.Background(), host)
	if err != nil {
		return nil, err
	}
	ips := make([]net.IP, 0, len(ipAddr))
	for _, addr := range ipAddr {
		ips = append(ips, addr.IP)
	}
	return ips, nil
}

func (c *Client) LookupIPv4(host string) ([]net.IP, error) {
	ips, err := c.LookupIP(host)
	if err != nil {
		return nil, err
	}
	var ipv4 []net.IP
	for _, ip := range ips {
		if len(ip) == net.IPv4len {
			ipv4 = append(ipv4, ip)
		}
	}
	return ipv4, nil
}

func (c *Client) LookupIPv6(host string) ([]net.IP, error) {
	ips, err := c.LookupIP(host)
	if err != nil {
		return nil, err
	}
	var ipv6 []net.IP
	for _, ip := range ips {
		if len(ip) == net.IPv6len {
			ipv6 = append(ipv6, ip)
		}
	}
	return ipv6, nil
}

func New() *Client {
	return &Client{
		resolver: net.Resolver{
			PreferGo: true,
		},
	}
}
