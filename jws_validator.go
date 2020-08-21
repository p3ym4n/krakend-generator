package generator

import "log"

func NewJWSValidator(alg, jwkUrl string) *JWSValidator {
	if stringInSlice(alg, JWSAlgorithms) {
		log.Fatalf("alg %q is not supported", alg)
	}
	return &JWSValidator{
		Alg:    alg,
		JwkUrl: jwkUrl,
	}
}

// JWSValidator is for making endpoints secure
type JWSValidator struct {
	Alg                string   `json:"alg"`
	JwkUrl             string   `json:"jwk-url"`
	Cache              bool     `json:"cache,omitempty"`
	CacheDuration      int      `json:"cache_duration,omitempty"`
	Audience           []string `json:"audience,omitempty"`
	RolesKey           string   `json:"roles_key,omitempty"`
	Roles              []string `json:"roles,omitempty"`
	Issuer             []string `json:"issuer,omitempty"`
	CookieKey          string   `json:"cookie_key,omitempty"`
	DisableJwkSecurity bool     `json:"disable_jwk_security,omitempty"`
	JwkFingerprints    []string `json:"jwk_fingerprints,omitempty"`
	CipherSuites       []int    `json:"cipher_suites,omitempty"`
	JwkLocalCa         string   `json:"jwk_local_ca,omitempty"`
}

// DefaultRoles will set the given roles authorized for all routes
func (v *JWSValidator) DefaultRoles(roles ...string) *JWSValidator {
	v.Roles = roles
	return v
}

// SetAlg will set the alg for JWSValidator
func (v *JWSValidator) SetAlg(alg string) *JWSValidator {
	if stringInSlice(alg, JWSAlgorithms) {
		log.Fatalf("alg %q is not supported", alg)
	}
	v.Alg = alg
	return v
}

// SetUrl will set the jwks url for JWSValidator
func (v *JWSValidator) SetUrl(url string) *JWSValidator {
	v.JwkUrl = url
	return v
}

// SetCache will set the cache for some minute to prevent hammering the server
func (v *JWSValidator) SetCache(duration int) *JWSValidator {
	v.Cache = true
	v.CacheDuration = duration
	return v
}

// ExportWithRoles will generate a map[string]interface{} based on the fields value
func (v *JWSValidator) ExportWithRoles(roles ...string) map[string]interface{} {
	out := map[string]interface{}{
		"alg":                  v.Alg,
		"jwk-url":              v.JwkUrl,
		"disable_jwk_security": v.DisableJwkSecurity,
	}
	if v.Cache != false {
		out["cache"] = v.Cache
	}
	if v.CacheDuration != 0 {
		out["cache_duration"] = v.CacheDuration
	}
	if v.Audience != nil {
		out["audience"] = v.Audience
	}
	if v.Audience != nil {
		out["audience"] = v.Audience
	}
	if v.RolesKey != "" {
		out["roles_key"] = v.RolesKey
	}
	if v.Roles != nil || roles != nil {
		out["roles"] = append(roles, v.Roles...)
	}
	if v.Issuer != nil {
		out["issuer"] = v.Issuer
	}
	if v.CookieKey != "" {
		out["cookie_key"] = v.CookieKey
	}
	if v.JwkFingerprints != nil {
		out["jwk_fingerprints"] = v.JwkFingerprints
	}
	if v.CipherSuites != nil {
		out["cipher_suites"] = v.CipherSuites
	}
	if v.JwkLocalCa != "" {
		out["jwk_local_ca"] = v.JwkLocalCa
	}
	return out
}
