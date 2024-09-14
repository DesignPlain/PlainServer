package cloudwatch

type LogDestinationPolicy struct {
	// The policy document. This is a JSON formatted string.
	AccessPolicy string `json:"accessPolicy,omitempty" yaml:"accessPolicy,omitempty"`

	// A name for the subscription filter
	DestinationName string `json:"destinationName,omitempty" yaml:"destinationName,omitempty"`

	// Specify true if you are updating an existing destination policy to grant permission to an organization ID instead of granting permission to individual AWS accounts.
	ForceUpdate bool `json:"forceUpdate,omitempty" yaml:"forceUpdate,omitempty"`
}
