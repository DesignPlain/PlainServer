package types

type Ssmincidents_ReplicationSetRegion struct {
	/*
	   The Amazon Resource name (ARN) of the customer managed key. If omitted, AWS manages the AWS KMS keys for you, using an AWS owned key, as indicated by a default value of `DefaultKey`.

	   The following arguments are optional:
	*/
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`

	// The name of the Region, such as `ap-southeast-2`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The current status of the Region.
	   - Valid Values: `ACTIVE` | `CREATING` | `UPDATING` | `DELETING` | `FAILED`
	*/
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// More information about the status of a Region.
	StatusMessage string `json:"statusMessage,omitempty" yaml:"statusMessage,omitempty"`
}
