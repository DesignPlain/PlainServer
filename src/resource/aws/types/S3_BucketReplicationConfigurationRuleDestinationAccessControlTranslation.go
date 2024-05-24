package types

type S3_BucketReplicationConfigurationRuleDestinationAccessControlTranslation struct {
	// The override value for the owner on replicated objects. Currently only `Destination` is supported.
	Owner string `json:"owner,omitempty" yaml:"owner,omitempty"`
}
