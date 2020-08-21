package generator

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

func Default() *App {
	return &App{
		Name:     "Api Gateway",
		Port:     8080,
		Timeout:  "10s",
		CacheTTL: "300s",
		ExtraConfig: ExtraConfig{
			"github_com/devopsfaith/krakend-cors": {
				"max_age":           "12h",
				"allow_credentials": true,
				"allow_origins":     []string{"http://localhost:8080"},
				"allow_methods":     []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
				"allow_headers":     []string{"*"},
				"expose_headers":    []string{"*"},
			},
			"github_com/devopsfaith/krakend-gologging": {
				"format": "logstash",
				"level":  "DEBUG",
				"prefix": "[KRAKEND]",
				"stdout": true,
				"syslog": true,
			},
			"github_com/devopsfaith/krakend-logstash": {
				"enabled": true,
			},
			"github_com/devopsfaith/krakend-metrics": {
				"listen_address": "8081",
			},
		},
	}
}

// App represent the root layer of krakend config which holds the application layer logic
type App struct {
	// Version your configuration
	Version int `json:"version"`
	// Name your configuration
	Name string `json:"name"`
	// a Port which krakend will listen to (this can be override by the cli)
	Port int `json:"port"`
	// global Timeout for all requests
	Timeout string `json:"timeout"`
	// global CacheTTL for all requests
	CacheTTL string `json:"cache_ttl"`
	// global ExtraConfig for all requests
	ExtraConfig ExtraConfig `json:"extra_config,omitempty"`
	// Endpoint which krakend must handle
	Endpoints []Endpoint `json:"endpoints"`
}

// AddEndpoints will add 1 or more endpoints into the app
func (app *App) AddEndpoints(endpoints ...*Endpoint) {
	for _, ep := range endpoints {
		app.Endpoints = append(app.Endpoints, *ep)
	}
}

// Generate wil generate the final result into the given path
func (app *App) Generate(path string) error {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(app)
	if err != nil {
		return err
	}
	indentBuffer := new(bytes.Buffer)
	if err := json.Indent(indentBuffer, buffer.Bytes(), "", "  "); err != nil {
		return err
	}
	if err = ioutil.WriteFile(path, indentBuffer.Bytes(), 0644); err != nil {
		return err
	}
	return nil
}
