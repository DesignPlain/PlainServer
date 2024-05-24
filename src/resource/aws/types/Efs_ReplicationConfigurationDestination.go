package types

type Efs_ReplicationConfigurationDestination struct {
	// The availability zone in which the replica should be created. If specified, the replica will be created with One Zone storage. If omitted, regional storage will be used.
	AvailabilityZoneName string `json:"availabilityZoneName,omitempty" yaml:"availabilityZoneName,omitempty"`

	// The ID of the destination file system for the replication. If no ID is provided, then EFS creates a new file system with the default settings.
	FileSystemId string `json:"fileSystemId,omitempty" yaml:"fileSystemId,omitempty"`

	// The Key ID, ARN, alias, or alias ARN of the KMS key that should be used to encrypt the replica file system. If omitted, the default KMS key for EFS `/aws/elasticfilesystem` will be used.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// The region in which the replica should be created.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	//
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
