//go:build stats
// +build stats

package main

import (
	"github.com/Jason-Stan-Lee/go-tun2socks/v2/common/stats/session"
)

func init() {
	addPostFlagsInitFn(func() {
		if *args.Stats {
			sessionStater = session.NewSimpleSessionStater()
			sessionStater.Start()
		} else {
			sessionStater = nil
		}
	})
}
