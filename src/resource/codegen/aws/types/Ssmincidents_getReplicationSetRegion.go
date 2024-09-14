package types

type Ssmincidents_getReplicationSetRegion struct {
	// More information about the status of a Region.
	StatusMessage string `json:"statusMessage,omitempty" yaml:"statusMessage,omitempty"`

	// The ARN of the AWS Key Management Service (AWS KMS) encryption key.
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`

	// The name of the Region.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The current status of the Region.
	   - Valid Values: `ACTIVE` | `CREATING` | `UPDATING` | `DELETING` | `FAILED`
	*/
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
