//go:build v2ray
// +build v2ray

package main

import (
	"context"
	"flag"
	"io/ioutil"
	"strings"

	vcore "github.com/v2fly/v2ray-core/v5"
	vproxyman "github.com/v2fly/v2ray-core/v5/app/proxyman"
	vbytespool "github.com/v2fly/v2ray-core/v5/common/bytespool"

	"github.com/Jason-Stan-Lee/go-tun2socks/v2/common/log"
	"github.com/Jason-Stan-Lee/go-tun2socks/v2/core"
	"github.com/Jason-Stan-Lee/go-tun2socks/v2/proxy/v2ray"
)

func init() {
	args.addFlag(fUdpTimeout)

	args.VConfig = flag.String("vconfig", "config.json", "Config file for v2ray, in JSON format, and note that routing in v2ray could not violate routes in the routing table")
	args.SniffingType = flag.String("sniffingType", "http,tls", "Enable domain sniffing for specific kind of traffic in v2ray")

	registerHandlerCreater("v2ray", func() {
		core.SetBufferPool(vbytespool.GetPool(core.BufSize))

		configBytes, err := ioutil.ReadFile(*args.VConfig)
		if err != nil {
			log.Fatalf("invalid vconfig file")
		}
		var validSniffings []string
		sniffings := strings.Split(*args.SniffingType, ",")
		for _, s := range sniffings {
			if s == "http" || s == "tls" {
				validSniffings = append(validSniffings, s)
			}
		}

		v, err := vcore.StartInstance("json", configBytes)
		if err != nil {
			log.Fatalf("start V instance failed: %v", err)
		}

		sniffingConfig := &vproxyman.SniffingConfig{
			Enabled:             true,
			DestinationOverride: validSniffings,
		}
		if len(validSniffings) == 0 {
			sniffingConfig.Enabled = false
		}

		ctx := vproxyman.ContextWithSniffingConfig(context.Background(), sniffingConfig)

		core.RegisterTCPConnHandler(v2ray.NewTCPHandler(ctx, v))
		core.RegisterUDPConnHandler(v2ray.NewUDPHandler(ctx, v, *args.UdpTimeout))
	})
}
