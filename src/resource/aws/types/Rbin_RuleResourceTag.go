package types

type Rbin_RuleResourceTag struct {
	/*
	   The tag key.

	   The following argument is optional:
	*/
	ResourceTagKey string `json:"resourceTagKey,omitempty" yaml:"resourceTagKey,omitempty"`

	// The tag value.
	ResourceTagValue string `json:"resourceTagValue,omitempty" yaml:"resourceTagValue,omitempty"`
}
