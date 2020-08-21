package generator

import (
	"log"
)

// NOOPEndpoint will make a proxy only Endpoint
func NOOPEndpoint(method, uri string, backends ...*Backend) *Endpoint {
	return NewEndpoint(method, uri, NOOP, backends...)
}

// JsonEndpoint will make an Endpoint with json encoding response
func JsonEndpoint(method, uri string, backends ...*Backend) *Endpoint {
	return NewEndpoint(method, uri, JSON, backends...)
}

// NegotiateEndpoint will make an Endpoint with encoding based on the Backend encoding
func NegotiateEndpoint(method, uri string, backends ...*Backend) *Endpoint {
	return NewEndpoint(method, uri, NEGOTIATE, backends...)
}

// StringEndpoint will make an Endpoint with string encoding response
func StringEndpoint(method, uri string, backends ...*Backend) *Endpoint {
	return NewEndpoint(method, uri, STRING, backends...)
}

// NewEndpoint will make an Endpoint with the given arguments
func NewEndpoint(method, uri, encoding string, backends ...*Backend) *Endpoint {
	if !stringInSlice(method, HttpMethods) {
		log.Fatalf("method %q is not supported", method)
	}
	if !stringInSlice(encoding, EndpointEncodings) {
		log.Fatalf("encoding %q is not supported", encoding)
	}
	backendValues := make([]Backend, 0)
	for _, backend := range backends {
		backendValues = append(backendValues, *backend)
	}
	headers := []string{
		"Accept-Encoding", "Host", "User-Agent", "X-Forwarded-For",
		"Content-Type", "Cache-Control", "Connection",
	}
	return &Endpoint{
		URI:      uri,
		Backend:  backendValues,
		Method:   method,
		Encoding: encoding,
		Headers:  headers,
	}
}

// Endpoint consist of one or more Backend s which defines a route in gateway
type Endpoint struct {
	URI string `json:"endpoint"`
	// Method is one of http.Method s
	Method string `json:"method"`
	// Encoding is one of EndpointEncodings
	Encoding string `json:"output_encoding"`
	// QueryStrings consist of query strings that will be passed to Backend by gateway
	// for passing all query strings you can use * or PassAllQueryStrings method.
	QueryStrings []string `json:"querystring_params,omitempty"`
	// Headers consist of all header that will be passed to Backend by gateway
	// by default these headers will be passed:
	// Accept-Encoding, Host, User-Agent, X-Forwarded-For
	// also by adding "Cookie" header all cookies will be passed to Backend
	// for passing all headers you can use * or PassAllHeaders method.
	Headers []string `json:"headers_to_pass,omitempty"`
	// ConcurrentCalls more than 1 can speeds up the response rate, but with
	// more load on the backend, for example 3 is a good number
	ConcurrentCalls uint                              `json:"concurrent_calls,omitempty"`
	Backend         []Backend                         `json:"backend"`
	ExtraConfig     map[string]map[string]interface{} `json:"extra_config,omitempty"`
}

// SetMethod will set the Endpoint method
func (ep *Endpoint) SetMethod(method string) *Endpoint {
	if stringInSlice(method, HttpMethods) {
		ep.Method = method
	} else {
		log.Fatalf("method %q is not supported", method)
	}
	return ep
}

// SetEncoding will set the Endpoint encoding
func (ep *Endpoint) SetEncoding(encoding string) *Endpoint {
	if stringInSlice(encoding, EndpointEncodings) {
		ep.Encoding = encoding
	} else {
		log.Fatalf("encoding %q is not supported", encoding)
	}
	return ep
}

// SetConcurrent will tell krakend to send more than one request to Backend to get faster response
func (ep *Endpoint) SetConcurrent(numberOfCalls uint) *Endpoint {
	if numberOfCalls > 1 {
		ep.ConcurrentCalls = numberOfCalls
	} else {
		log.Fatal("number of concurrent calls must be more than 1")
	}
	return ep
}

// RequestsAreSequential will declare that the Backend s should be called after each other
func (ep *Endpoint) RequestsAreSequential() *Endpoint {
	ep.SetConfig("github.com/devopsfaith/krakend/proxy", map[string]interface{}{
		"sequential": true,
	})
	return ep
}

// PassQueryString will pass the client specified query_string to Backend
func (ep *Endpoint) PassQueryString(queryStrings ...string) *Endpoint {
	ep.QueryStrings = append(ep.QueryStrings, queryStrings...)
	return ep
}

// PassAllQueryStrings will pass all client query_strings to Backend
func (ep *Endpoint) PassAllQueryStrings() *Endpoint {
	ep.QueryStrings = []string{"*"}
	return ep
}

// PassHeader will pass the specified client header to Backend
func (ep *Endpoint) PassHeader(headers ...string) *Endpoint {
	ep.Headers = append(ep.Headers, headers...)
	return ep
}

// PassAllHeaders will pass all client headers to Backend
func (ep *Endpoint) PassAllHeaders() *Endpoint {
	ep.Headers = []string{"*"}
	return ep
}

// with SetConfig you can explicitly set an extra_config for Endpoint
func (ep *Endpoint) SetConfig(name string, config map[string]interface{}) *Endpoint {
	if ep.ExtraConfig == nil {
		ep.ExtraConfig = ExtraConfig{name: config}
	} else {
		ep.ExtraConfig[name] = config
	}
	return ep
}

// AddBackend will add a Backend to the Endpoint
func (ep *Endpoint) AddBackend(backends ...Backend) *Endpoint {
	ep.Backend = append(ep.Backend, backends...)
	return ep
}

// Authenticate will authenticate the user request based on jwk defined in JWSValidator and roles
func (ep *Endpoint) Authenticate(v JWSValidator, roles ...string) *Endpoint {
	ep.SetConfig("github.com/devopsfaith/krakend-jose/validator", v.ExportWithRoles(roles...))
	return ep
}
