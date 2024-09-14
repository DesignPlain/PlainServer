package types

type Kms_getKeyMultiRegionConfiguration struct {
	// Indicates whether the KMS key is a `PRIMARY` or `REPLICA` key.
	MultiRegionKeyType string `json:"multiRegionKeyType,omitempty" yaml:"multiRegionKeyType,omitempty"`

	// The key ARN and Region of the primary key. This is the current KMS key if it is the primary key.
	PrimaryKeys []Kms_getKeyMultiRegionConfigurationPrimaryKey `json:"primaryKeys,omitempty" yaml:"primaryKeys,omitempty"`

	// The key ARNs and Regions of all replica keys. Includes the current KMS key if it is a replica key.
	ReplicaKeys []Kms_getKeyMultiRegionConfigurationReplicaKey `json:"replicaKeys,omitempty" yaml:"replicaKeys,omitempty"`
}
