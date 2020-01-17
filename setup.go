package delay

import (
	"strconv"

	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/plugin/metrics"

	"github.com/caddyserver/caddy"
)

// init registers this plugin.
func init() { plugin.Register("delay", setup) }

// setup is the function that gets called when the config parser sees the token "delay". Setup is responsible
// for parsing any extra options the delay plugin may have. The first token this function sees is "delay".
func setup(c *caddy.Controller) (*Delay, error) {
	d := &Delay{Delay: 2000}
	c.Next() // Ignore "delay" and give us the next token.
	args := c.RemainingArgs()
	if len(args) > 1 {
		return nil, c.ArgErr()
	}

	if len(args) == 0 {
		continue
	}

	delay, err := strconv.ParseInt(args[0], 10, 32)
	if err != nil {
		return nil, err
	}
	d.Delay = uint64(delay)
	return d, nil

	// Add a startup function that will -- after all plugins have been loaded -- check if the
	// prometheus plugin has been used - if so we will export metrics. We can only register
	// this metric once, hence the "once.Do".
	c.OnStartup(func() error {
		once.Do(func() { metrics.MustRegister(c, requestCount) })
		return nil
	})

	// Add the Plugin to CoreDNS, so Servers can use it in their plugin chain.
	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return Delay{Next: next}
	})

	// All OK, return a nil error.
	return nil
}
