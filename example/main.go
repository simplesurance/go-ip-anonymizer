package main

import (
	"log"

	goanonymizer "github.com/simplesurance/go-ip-anonymizer"
)

func main() {
	ip4 := "192.168.1.12"
	ip6 := "bbd1:e95a:adbb:b29a:e38b:577f:6f9a:1fa7"
	newIP4, err := goanonymizer.NewAnonymize("", "").AnonymizeIp(ip4)
	if err != nil {
		log.Println(err)
	}
	newIP6, err := goanonymizer.NewAnonymize("", "").AnonymizeIp(ip6)
	if err != nil {
		log.Println(err)
	}

	log.Printf("%s is transformed to %s", ip4, newIP4) //192.168.1.12 is transformed to 192.168.1.0
	log.Printf("%s is transformed to %s", ip6, newIP6) //bbd1:e95a:adbb:b29a:e38b:577f:6f9a:1fa7 is transformed to bbd1:e95a:adbb:b29a::

}
