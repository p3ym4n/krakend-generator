package generator

type ExtraConfig map[string]map[string]interface{}

const (
	// backend & endpoint accepted encodings
	JSON   = "json"
	XML    = "xml"
	RSS    = "rss"
	STRING = "string"
	NOOP   = "no-op"
	// only for endpoint
	NEGOTIATE = "negotiate"
	YAML      = "yaml"
)

const (
	GET     = "GET"
	HEAD    = "HEAD"
	POST    = "POST"
	PUT     = "PUT"
	PATCH   = "PATCH"
	DELETE  = "DELETE"
	CONNECT = "CONNECT"
	OPTIONS = "OPTIONS"
	TRACE   = "TRACE"
)

var (
	BackendEncodings  = []string{JSON, XML, RSS, STRING, NOOP}
	EndpointEncodings = []string{JSON, NEGOTIATE, STRING, NOOP}
	HttpMethods       = []string{GET, HEAD, OPTIONS, POST, PUT, PATCH, DELETE, CONNECT, TRACE}
	JWSAlgorithms     = []string{
		"EdDSA", "HS256", "HS384", "HS512", "RS256", "RS384", "RS512",
		"ES256", "ES384", "ES512", "PS256", "PS384", "PS512",
	}
)
