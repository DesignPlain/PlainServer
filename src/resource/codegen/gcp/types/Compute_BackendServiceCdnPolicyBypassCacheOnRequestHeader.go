package types

type Compute_BackendServiceCdnPolicyBypassCacheOnRequestHeader struct {
	// The header field name to match on when bypassing cache. Values are case-insensitive.
	HeaderName string `json:"headerName,omitempty" yaml:"headerName,omitempty"`
}
