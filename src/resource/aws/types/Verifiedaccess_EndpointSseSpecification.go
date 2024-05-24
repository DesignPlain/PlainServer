package types

type Verifiedaccess_EndpointSseSpecification struct {
	//
	CustomerManagedKeyEnabled bool `json:"customerManagedKeyEnabled,omitempty" yaml:"customerManagedKeyEnabled,omitempty"`

	//
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`
}
