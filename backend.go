package generator

import (
	"log"
)

func NOOPBackend(method, host, uri string) *Backend {
	return NewBackend(method, host, uri, NOOP)
}

func JsonBackend(method, host, uri string) *Backend {
	return NewBackend(method, host, uri, JSON)
}

func XMLBackend(method, host, uri string) *Backend {
	return NewBackend(method, host, uri, XML)
}

func RSSBackend(method, host, uri string) *Backend {
	return NewBackend(method, host, uri, RSS)
}

func NewBackend(method, host, uri, encoding string) *Backend {
	if !stringInSlice(method, HttpMethods) {
		log.Fatalf("method %q is not supported", method)
	}
	if !stringInSlice(encoding, BackendEncodings) {
		log.Fatalf("encoding %q is not supported", encoding)
	}
	return &Backend{
		URI:      uri,
		Method:   method,
		Host:     []string{host},
		Encoding: encoding,
	}
}

type Backend struct {
	URI      string   `json:"url_pattern"`
	Method   string   `json:"method,omitempty"`
	Encoding string   `json:"encoding,omitempty"`
	Host     []string `json:"host"`
	// filters only the items in blacklist
	BlackList []string `json:"blacklist,omitempty"`
	// filters only the items in whitelist
	Whitelist []string `json:"whitelist,omitempty"`
	Group     string   `json:"group,omitempty"`
	// this will rename the items
	Mapping map[string]string `json:"mapping,omitempty"`
	// will capture only that section of the result
	Target string `json:"target,omitempty"`

	// when the root level of response is an array or a collection
	// by default they will be added as the value of "collection" key
	// which that name can be changed by use of mapping
	IsCollection bool `json:"is_collection,omitempty"`

	ExtraConfig map[string]map[string]interface{} `json:"extra_config,omitempty"`

	// modifiers are for martian plugin
	modifiers map[string]map[string]interface{}
}

// SetEncoding will set the Backend Encoding
func (b *Backend) SetEncoding(encoding string) *Backend {
	if stringInSlice(encoding, BackendEncodings) {
		b.Encoding = encoding
	} else {
		log.Fatalf("encoding %q is not supported", encoding)
	}
	return b
}

// SetMethod will set the Backend Method
func (b *Backend) SetMethod(method string) *Backend {
	if stringInSlice(method, HttpMethods) {
		b.Method = method
	} else {
		log.Fatalf("method %q is not supported", method)
	}
	return b
}

// AddHost will define one or more Host s
func (b *Backend) AddHost(hosts ...string) *Backend {
	b.Host = append(b.Host, hosts...)
	return b
}

// AddMapping renames an item in Backend response
func (b *Backend) AddMapping(from, to string) *Backend {
	if b.Mapping == nil {
		b.Mapping = map[string]string{from: to}
	} else {
		b.Mapping[from] = to
	}
	return b
}

// SetBlacklist delete one or more items from Backend response
func (b *Backend) SetBlacklist(items ...string) *Backend {
	b.BlackList = items
	return b
}

// SetWhitelist only returns one or more items from Backend response
func (b *Backend) SetWhitelist(items ...string) *Backend {
	b.Whitelist = items
	return b
}

// SetConfig allows you to explicitly add an extra_config into the Backend
func (b *Backend) SetConfig(name string, config map[string]interface{}) *Backend {
	if b.ExtraConfig == nil {
		b.ExtraConfig = ExtraConfig{name: config}
	} else {
		b.ExtraConfig[name] = config
	}
	return b
}

// EnableCache will cache the Backend responses and will heavily increase the memory usage!
func (b *Backend) EnableCache() *Backend {
	return b.SetConfig("github.com/devopsfaith/krakend-httpcache", map[string]interface{}{})
}

// ShadowEnabled will show all the requests into the Backend but will not return its response
func (b *Backend) ShadowEnabled() *Backend {
	return b.SetConfig("github.com/devopsfaith/krakend/proxy", map[string]interface{}{
		"shadow": true,
	})
}

// CircuitBreaker will enable removing the failed Host s from the list
func (b *Backend) CircuitBreaker(interval, timeout, maxErrors int, logChanges bool) *Backend {
	return b.SetConfig("github.com/devopsfaith/krakend-circuitbreaker/gobreaker", map[string]interface{}{
		"interval":        interval,
		"timeout":         timeout,
		"maxErrors":       maxErrors,
		"logStatusChange": logChanges,
	})
}
