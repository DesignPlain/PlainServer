package types

type Networkconnectivity_ServiceConnectionPolicyPscConnectionErrorInfo struct {
	// The logical grouping to which the "reason" belongs.
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`

	// Additional structured details about this error.
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// The reason of the error.
	Reason string `json:"reason,omitempty" yaml:"reason,omitempty"`
}
