package cloudfront

type KeyGroup struct {
	// A comment to describe the key group..
	Comment string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// A list of the identifiers of the public keys in the key group.
	Items []string `json:"items,omitempty" yaml:"items,omitempty"`

	// A name to identify the key group.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
