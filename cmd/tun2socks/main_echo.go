//go:build echo
// +build echo

package main

import (
	"github.com/Jason-Stan-Lee/go-tun2socks/v2/core"
	"github.com/Jason-Stan-Lee/go-tun2socks/v2/proxy/echo"
)

func init() {
	registerHandlerCreater("echo", func() {
		core.RegisterTCPConnHandler(echo.NewTCPHandler())
		core.RegisterUDPConnHandler(echo.NewUDPHandler())
	})
}
