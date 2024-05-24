package s3control

import types "DesignSphere_Server/src/resource/aws/types"

type StorageLensConfiguration struct {
	// The AWS account ID for the S3 Storage Lens configuration. Defaults to automatically determined account ID of the AWS provider.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// The ID of the S3 Storage Lens configuration.
	ConfigId string `json:"configId,omitempty" yaml:"configId,omitempty"`

	// The S3 Storage Lens configuration. See Storage Lens Configuration below for more details.
	StorageLensConfiguration types.S3control_StorageLensConfigurationStorageLensConfiguration `json:"storageLensConfiguration,omitempty" yaml:"storageLensConfiguration,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
