package goanonymizer

import (
	"errors"
	"net"
)

const DefaultIPv4Mask = "255.255.255.0"
const DefaultIPv6Mask = ""ffff:ffff:ffff:ffff:0000:0000:0000:0000""
//Anonymize struct that hold mask for V4 and V6.
type Anonymize struct {
	v4Mask net.IPMask
	v6Mask net.IPMask
}

//NewAnonymize returns instance of Anonymize.
//If v4Mask and v6Mask are not provided, default values are used.
func NewAnonymize(v4Mask, v6Mask string) *Anonymize {
	ipV4 := net.ParseIP(v4Mask)
	if ipV4 == nil {
		ipV4 = net.ParseIP("255.255.255.0")
	}
	
	ipV6 := net.ParseIP(v6Mask)
	if ipV6 == nil {
		ipV6 = net.ParseIP("ffff:ffff:ffff:ffff:0000:0000:0000:0000")
	}

	return &Anonymize{v4Mask: net.IPMask(ipV4.To4()), v6Mask: net.IPMask(ipV6)}
}

//AnonymizeIp gets IP and transform it based on v4Mask and v6Mask.
func (a Anonymize) AnonymizeIp(ipStr string) (string, error) {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return "", errors.New("IP is not valid")
	}
	
	if ipv4 := ip.To4(); ipv4 != nil {
		return ipv4.Mask(a.v4Mask).String(), nil
	}

	return ip.To16().Mask(a.v6Mask).String(), nil
}
