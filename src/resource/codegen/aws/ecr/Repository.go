package ecr

import types "libds/aws/types"

type Repository struct {
	// Name of the repository.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Encryption configuration for the repository. See below for schema.
	EncryptionConfigurations []types.Ecr_RepositoryEncryptionConfiguration `json:"encryptionConfigurations,omitempty" yaml:"encryptionConfigurations,omitempty"`

	/*
	   If `true`, will delete the repository even if it contains images.
	   Defaults to `false`.
	*/
	ForceDelete bool `json:"forceDelete,omitempty" yaml:"forceDelete,omitempty"`

	// Configuration block that defines image scanning configuration for the repository. By default, image scanning must be manually triggered. See the [ECR User Guide](https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-scanning.html) for more information about image scanning.
	ImageScanningConfiguration types.Ecr_RepositoryImageScanningConfiguration `json:"imageScanningConfiguration,omitempty" yaml:"imageScanningConfiguration,omitempty"`

	// The tag mutability setting for the repository. Must be one of: `MUTABLE` or `IMMUTABLE`. Defaults to `MUTABLE`.
	ImageTagMutability string `json:"imageTagMutability,omitempty" yaml:"imageTagMutability,omitempty"`
}
