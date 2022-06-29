// Package ipranges reads GCP compute IP ranges.
//
// Google Cloud publishes a JSON-formatted list of customer-usable
// global and regional external IP address ranges here:
// https://www.gstatic.com/ipranges/cloud.json.
package ipranges

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
)

type IPRanges struct {
	SyncToken    string
	CreationTime string
	Prefixes     []struct {
		IPV4Prefix string
		IPV6Prefix string
		Service    string
		Scope      string
	}
}

// Regions returns a list of regions matching the given prefix.
func (r *IPRanges) Regions(prefix string) []string {
	m := make(map[string]struct{})
	for _, p := range r.Prefixes {
		if strings.HasPrefix(p.Scope, prefix) {
			m[p.Scope] = struct{}{}
		}
	}
	return orderedKeys(m)
}

func orderedKeys(m map[string]struct{}) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func (r *IPRanges) Services() []string {
	m := make(map[string]struct{})
	for _, p := range r.Prefixes {
		m[p.Service] = struct{}{}
	}
	return orderedKeys(m)
}

// Fetch the IP ranges data from https://www.gstatic.com/ipranges/cloud.json.
func Fetch(ctx context.Context) (*IPRanges, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.gstatic.com/ipranges/cloud.json", strings.NewReader(""))
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not fetch data: status = %d (%v)", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	out := &IPRanges{}
	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return nil, err
	}
	return out, nil
}
