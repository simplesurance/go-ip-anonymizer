# IP Anonymizer for Golang

This Golang package provides functionality to anonymize IP-Addresses.
IP-Addresses are anonymities by zeroing the host-part of an address.

The host part of addresses that are anonymities be specified by providing a
subnet mask.

## Example

```go
package main

import (
    "fmt"
    "net"
    "os"

    "github.com/simplesurance/go-ip-anonymizer/ipanonymizer"
)

func main() {
    const ip4 = "192.168.1.12"
    const ip6 = "bbd1:e95a:adbb:b29a:e38b:577f:6f9a:1fa7"

    // Create an anonymizer with a /16 IPv6 subnet mask and
    // a /64 IPv6 // subnet mask.
    anonymizer := ipanonymizer.NewWithMask(
        net.CIDRMask(16, 32),
        net.CIDRMask(64, 128),
    )

    anonIP4, err := anonymizer.IPString(ip4)
    exitOnErr(err)
    fmt.Printf("%s anonymized to %s\n", ip4, anonIP4)
    // Prints: 192.168.1.12 anonymized to 192.168.0.0

    anonIP6, err := anonymizer.IPString(ip6)
    exitOnErr(err)
    fmt.Printf("%s anonymized to %s\n", ip6, anonIP6)
    // Prints: bbd1:e95a:adbb:b29a:e38b:577f:6f9a:1fa7 anonymized to bbd1:e95a:adbb:b29a::
}
```
