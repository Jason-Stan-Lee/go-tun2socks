//go:build dnsfallback
// +build dnsfallback

package main

import (
	"flag"

	"github.com/Jason-Stan-Lee/go-tun2socks/v2/core"
	"github.com/Jason-Stan-Lee/go-tun2socks/v2/proxy/dnsfallback"
)

func init() {
	args.DnsFallback = flag.Bool("dnsFallback", false, "Enable DNS fallback over TCP (overrides the UDP proxy handler).")

	registerHandlerCreater("dnsfallback", func() {
		core.RegisterUDPConnHandler(dnsfallback.NewUDPHandler())
	})
}
