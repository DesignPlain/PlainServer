package types

type Verifiedaccess_GroupSseConfiguration struct {
	// ARN of the KMS key to use.
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`

	//
	CustomerManagedKeyEnabled bool `json:"customerManagedKeyEnabled,omitempty" yaml:"customerManagedKeyEnabled,omitempty"`
}
