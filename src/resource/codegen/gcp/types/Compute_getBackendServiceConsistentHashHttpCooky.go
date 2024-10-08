package types

type Compute_getBackendServiceConsistentHashHttpCooky struct {
	/*
	   The name of the Backend Service.

	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Path to set for the cookie.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// Lifetime of the cookie.
	Ttls []Compute_getBackendServiceConsistentHashHttpCookyTtl `json:"ttls,omitempty" yaml:"ttls,omitempty"`
}
