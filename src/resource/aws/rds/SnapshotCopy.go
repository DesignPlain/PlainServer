package rds

type SnapshotCopy struct {
	// The Destination region to place snapshot copy.
	DestinationRegion string `json:"destinationRegion,omitempty" yaml:"destinationRegion,omitempty"`

	// KMS key ID.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// The name of an option group to associate with the copy of the snapshot.
	OptionGroupName string `json:"optionGroupName,omitempty" yaml:"optionGroupName,omitempty"`

	// The Identifier for the snapshot.
	TargetDbSnapshotIdentifier string `json:"targetDbSnapshotIdentifier,omitempty" yaml:"targetDbSnapshotIdentifier,omitempty"`

	// Whether to copy existing tags. Defaults to `false`.
	CopyTags bool `json:"copyTags,omitempty" yaml:"copyTags,omitempty"`

	// he URL that contains a Signature Version 4 signed request.
	PresignedUrl string `json:"presignedUrl,omitempty" yaml:"presignedUrl,omitempty"`

	// Snapshot identifier of the source snapshot.
	SourceDbSnapshotIdentifier string `json:"sourceDbSnapshotIdentifier,omitempty" yaml:"sourceDbSnapshotIdentifier,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The external custom Availability Zone.
	TargetCustomAvailabilityZone string `json:"targetCustomAvailabilityZone,omitempty" yaml:"targetCustomAvailabilityZone,omitempty"`
}
