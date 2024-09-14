package types

type Compute_RegionUrlMapTest struct {
	// Host portion of the URL.
	Host string `json:"host,omitempty" yaml:"host,omitempty"`

	// Path portion of the URL.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// A reference to expected RegionBackendService resource the given URL should be mapped to.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`

	// Description of this test case.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
