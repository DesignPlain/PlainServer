package types

type Compute_BackendServiceConsistentHashHttpCookie struct {
	// Name of the cookie.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Path to set for the cookie.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	/*
	   Lifetime of the cookie.
	   Structure is documented below.
	*/
	Ttl Compute_BackendServiceConsistentHashHttpCookieTtl `json:"ttl,omitempty" yaml:"ttl,omitempty"`
}
