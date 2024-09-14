package types

type Ssoadmin_InstanceAccessControlAttributesAttribute struct {
	// The name of the attribute associated with your identities in your identity source. This is used to map a specified attribute in your identity source with an attribute in AWS SSO.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// The value used for mapping a specified attribute to an identity source. See AccessControlAttributeValue
	Values []Ssoadmin_InstanceAccessControlAttributesAttributeValue `json:"values,omitempty" yaml:"values,omitempty"`
}
