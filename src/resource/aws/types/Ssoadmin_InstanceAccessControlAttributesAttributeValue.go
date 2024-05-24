package types

type Ssoadmin_InstanceAccessControlAttributesAttributeValue struct {
	// The identity source to use when mapping a specified attribute to AWS SSO.
	Sources []string `json:"sources,omitempty" yaml:"sources,omitempty"`
}
