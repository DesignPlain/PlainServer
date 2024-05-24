package types

type Kms_getKeyMultiRegionConfigurationReplicaKey struct {
	// The key ARN of a primary or replica key of a multi-Region key.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// The AWS Region of a primary or replica key in a multi-Region key.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
}
