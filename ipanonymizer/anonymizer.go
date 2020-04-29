// Package ipanonymizer provides functionality to anonymize IP-Addresses.
package ipanonymizer

import (
	"errors"
	"net"
)

var (
	defaultIPv4Mask = net.IPv4Mask(255, 255, 255, 0)                                                     // /24
	defaultIPv6Mask = net.IPMask{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0, 0, 0, 0, 0, 0, 0, 0} // /64
)

// Anonymizer anonymizes IP-Addresses by zeroing their host part.
type Anonymizer struct {
	v4Mask net.IPMask
	v6Mask net.IPMask
}

// NewWithMask returns an IPAnonymizer instance with custom subnet masks.
// The v4SubnetMask v6SubnetMask define the host part of IP addresses that are
// zeroed.
// For example when using a /16 v4SubnetMask, the IP-Address 8.8.8.8 would be
// anonymized to  8.8.0.0.
func NewWithMask(v4SubnetMask, v6SubnetMask net.IPMask) *Anonymizer {
	return &Anonymizer{
		v4Mask: v4SubnetMask,
		v6Mask: v6SubnetMask,
	}
}

// New returns an IP-Anonymizer that uses the following subnet masks:
// IPv4: /24
// IPv6: /64
func New() *Anonymizer {
	return &Anonymizer{
		v4Mask: defaultIPv4Mask,
		v6Mask: defaultIPv6Mask,
	}
}

// IPv4 anonymizes an IPv4 address by zeroing it's host part.
func (a *Anonymizer) IPv4(ip net.IP) net.IP {
	return ip.Mask(a.v4Mask)
}

// IPv6 anonymizes an IPv4 address by zeroing it's host part.
func (a *Anonymizer) IPv6(ip net.IP) net.IP {
	return ip.Mask(a.v6Mask)
}

type ipVersion int

const (
	ipVersionUndefined ipVersion = iota
	ipv4                         = iota
	ipv6                         = iota
)

func ipVer(ipAddress string) ipVersion {
	// copied from net.ParseIP()
	for i := 0; i < len(ipAddress); i++ {
		switch ipAddress[i] {
		case '.':
			return ipv4
		case ':':
			return ipv6
		}
	}

	return ipVersionUndefined
}

// IPString anonymizes an IP address by zeroing it's host part.
// ipAddress must be in IPv4 or IPv6 notation.
func (a *Anonymizer) IPString(ipAddress string) (string, error) {
	switch ipVer(ipAddress) {
	case ipv4:
		ip4 := net.ParseIP(ipAddress)
		if ip4 == nil {
			return "", errors.New("invalid IPv4 address")
		}

		return a.IPv4(ip4).String(), nil

	case ipv6:
		ipv6 := net.ParseIP(ipAddress)
		if ipv6 == nil {
			return "", errors.New("invalid IPv6 address")
		}

		return a.IPv6(ipv6).String(), nil

	default:
		return "", errors.New("invalid ipv4 and ipv6 address")
	}
}
