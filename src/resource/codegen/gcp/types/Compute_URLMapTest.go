package types

type Compute_URLMapTest struct {
	// Description of this test case.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Host portion of the URL.
	Host string `json:"host,omitempty" yaml:"host,omitempty"`

	// Path portion of the URL.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// The backend service or backend bucket link that should be matched by this test.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`
}
