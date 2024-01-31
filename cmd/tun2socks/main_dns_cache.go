//go:build dnscache
// +build dnscache

package main

import (
	"flag"

	"github.com/Jason-Stan-Lee/go-tun2socks/v2/common/dns/cache"
)

func init() {
	args.DisableDnsCache = flag.Bool("disableDNSCache", false, "Disable DNS cache")

	addPostFlagsInitFn(func() {
		if *args.DisableDnsCache {
			dnsCache = nil
		} else {
			dnsCache = cache.NewSimpleDnsCache()
		}
	})
}
