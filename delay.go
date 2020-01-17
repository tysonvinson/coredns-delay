// Package delay is a CoreDNS plugin that sleeps for a configurable interval before passing to the next plugin
//
package delay

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/plugin/metrics"
	clog "github.com/coredns/coredns/plugin/pkg/log"

	"github.com/miekg/dns"
)

// Define log to be a logger with the plugin name in it. This way we can just use log.Info and
// friends to log.
var log = clog.NewWithPlugin("delay")

type Delay struct {
	Next plugin.Handler
}

// ServeDNS implements the plugin.Handler interface. This method gets called when delay is used
// in a Server.
func (e Delay) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {

	// Debug log that we've have seen the query. This will only be shown when the debug plugin is loaded.
	log.Debug("Received response")

	// Pause execution for configured interval
	time.Sleep(2000 * time.Millisecond)

	// Call next plugin (if any).
	return plugin.NextOrFailure(e.Name(), e.Next, ctx, pw, r)
}

// Name implements the Handler interface.
func (e Delay) Name() string { return "delay" }
