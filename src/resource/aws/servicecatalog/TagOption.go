package servicecatalog

type TagOption struct {
	// Tag option key.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	/*
	   Tag option value.

	   The following arguments are optional:
	*/
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// Whether tag option is active. Default is `true`.
	Active bool `json:"active,omitempty" yaml:"active,omitempty"`
}
