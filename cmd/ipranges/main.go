// The ipgranges tool fetches IP ranges used by GCP compute.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/dhowden/ipranges"
)

var (
	regions      = flag.Bool("regions", false, "only list regions")
	ipv6         = flag.Bool("ipv6", true, "show IPv6 addresses")
	ipv4         = flag.Bool("ipv4", true, "show IPv4 addresses")
	regionPrefix = flag.String("region", "", "prefix of region name to match")
	fetchTimeout = flag.Duration("fetch-timeout", 5*time.Second, "fetch timeout")
)

func main() {
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), *fetchTimeout)
	defer cancel()

	ir, err := ipranges.Fetch(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	if *regions {
		rs := ir.Regions(*regionPrefix)
		for _, r := range rs {
			fmt.Println(r)
		}
		return
	}

	out := make([]string, 0, len(ir.Prefixes))
	for _, p := range ir.Prefixes {
		if strings.HasPrefix(p.Scope, *regionPrefix) {
			if *ipv4 && p.IPV4Prefix != "" {
				out = append(out, p.IPV4Prefix)
			}
			if *ipv6 && p.IPV6Prefix != "" {
				out = append(out, p.IPV6Prefix)
			}
		}
	}

	sort.Strings(out)
	for _, x := range out {
		fmt.Println(x)
	}
}
