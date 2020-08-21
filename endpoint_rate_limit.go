package generator

// RateLimitGlobal sets a limit for all requests per second on the Endpoint
func (ep *Endpoint) RateLimitGlobal(max int) *Endpoint {
	ep.SetConfig("github.com/devopsfaith/krakend-ratelimit/juju/router", map[string]interface{}{
		"maxRate": max,
	})
	return ep
}

// RateLimitByIp sets a limit for requests for each client (identified by ip) per second on the Endpoint
func (ep *Endpoint) RateLimitByIp(max int) *Endpoint {
	ep.SetConfig("github.com/devopsfaith/krakend-ratelimit/juju/router", map[string]interface{}{
		"clientMaxRate": max,
		"strategy":      "ip",
	})
	return ep
}

// RateLimitByIp sets a limit for requests for each client (identified by header) per second on the Endpoint
func (ep *Endpoint) RateLimitByHeader(header string, max int) *Endpoint {
	ep.SetConfig("github.com/devopsfaith/krakend-ratelimit/juju/router", map[string]interface{}{
		"clientMaxRate": max,
		"strategy":      "header",
		"key":           header,
	})
	return ep
}

// RateLimitByIp sets a limit for all requests and for each client (identified by ip) per second on the Endpoint
func (ep *Endpoint) RateLimitByIpAndGlobal(clientMax, globalMax int) *Endpoint {
	ep.SetConfig("github.com/devopsfaith/krakend-ratelimit/juju/router", map[string]interface{}{
		"clientMaxRate": clientMax,
		"maxRate":       globalMax,
		"strategy":      "ip",
	})
	return ep
}

// RateLimitByIp sets a limit for all requests for each client (identified by header) per second on the Endpoint
func (ep *Endpoint) RateLimitByHeaderAndGlobal(header string, clientMax, globalMax int) *Endpoint {
	ep.SetConfig("github.com/devopsfaith/krakend-ratelimit/juju/router", map[string]interface{}{
		"clientMaxRate": clientMax,
		"maxRate":       globalMax,
		"strategy":      "header",
		"key":           header,
	})
	return ep
}

