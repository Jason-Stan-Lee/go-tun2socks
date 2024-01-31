//go:build redirect
// +build redirect

package main

import (
	"github.com/Jason-Stan-Lee/go-tun2socks/v2/core"
	"github.com/Jason-Stan-Lee/go-tun2socks/v2/proxy/redirect"
)

func init() {
	args.addFlag(fProxyServer)
	args.addFlag(fUdpTimeout)

	registerHandlerCreater("redirect", func() {
		core.RegisterTCPConnHandler(redirect.NewTCPHandler(*args.ProxyServer))
		core.RegisterUDPConnHandler(redirect.NewUDPHandler(*args.ProxyServer, *args.UdpTimeout))
	})
}
