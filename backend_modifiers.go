package generator

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func (b *Backend) addModifier(name string, config map[string]interface{}) *Backend {

	if b.modifiers == nil {
		b.modifiers = map[string]map[string]interface{}{name: config}
		b.SetConfig("github.com/devopsfaith/krakend-martian", map[string]interface{}{name: config})
	} else {
		b.modifiers[name] = config
		b.SetConfig("github.com/devopsfaith/krakend-martian", map[string]interface{}{
			"fifo.Group": map[string]interface{}{
				"scope":           []string{"request", "response"},
				"aggregateErrors": true,
				"modifiers":       b.modifiers,
			},
		})
	}

	return b
}

// RemoveResponseHeaders will remove one or more headers from Backend response
func (b *Backend) RemoveResponseHeaders(name ...string) *Backend {
	b.addModifier("header.Blacklist", map[string]interface{}{
		"scope": []string{"response"},
		"names": name,
	})
	return b
}

// InjectHeader will inject a header into the request when calling Backend
func (b *Backend) InjectHeader(name, value string, alsoToResponse bool) *Backend {
	scope := []string{"request"}
	if alsoToResponse {
		scope = append(scope, "response")
	}
	b.addModifier("header.Modifier", map[string]interface{}{
		"scope": scope,
		"name":  name,
		"value": value,
	})
	return b
}

// BearerAuth will add a bearer header when sending requests to Backend
func (b *Backend) BearerAuth(token string) *Backend {
	token = strings.Replace(token, "Bearer ", "", 1)
	b.InjectHeader("Authorization", fmt.Sprintf("Bearer %s", token), false)
	return b
}

// BasicAuth will add a basic auth header when sending requests to Backend
func (b *Backend) BasicAuth(username, password string) *Backend {
	token := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	b.InjectHeader("Authorization", fmt.Sprintf("Basic %s", token), false)
	return b
}


