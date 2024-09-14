package types

type Ecs_ClusterConfigurationManagedStorageConfiguration struct {
	// AWS Key Management Service key ID for the Fargate ephemeral storage.
	FargateEphemeralStorageKmsKeyId string `json:"fargateEphemeralStorageKmsKeyId,omitempty" yaml:"fargateEphemeralStorageKmsKeyId,omitempty"`

	// AWS Key Management Service key ID to encrypt the managed storage.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`
}
