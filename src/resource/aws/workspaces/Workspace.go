package workspaces

import types "DesignSphere_Server/src/resource/aws/types"

type Workspace struct {
	// Indicates whether the data stored on the user volume is encrypted.
	UserVolumeEncryptionEnabled bool `json:"userVolumeEncryptionEnabled,omitempty" yaml:"userVolumeEncryptionEnabled,omitempty"`

	// The ARN of a symmetric AWS KMS customer master key (CMK) used to encrypt data stored on your WorkSpace. Amazon WorkSpaces does not support asymmetric CMKs.
	VolumeEncryptionKey string `json:"volumeEncryptionKey,omitempty" yaml:"volumeEncryptionKey,omitempty"`

	// The WorkSpace properties.
	WorkspaceProperties types.Workspaces_WorkspaceWorkspaceProperties `json:"workspaceProperties,omitempty" yaml:"workspaceProperties,omitempty"`

	// The ID of the bundle for the WorkSpace.
	BundleId string `json:"bundleId,omitempty" yaml:"bundleId,omitempty"`

	// The ID of the directory for the WorkSpace.
	DirectoryId string `json:"directoryId,omitempty" yaml:"directoryId,omitempty"`

	// Indicates whether the data stored on the root volume is encrypted.
	RootVolumeEncryptionEnabled bool `json:"rootVolumeEncryptionEnabled,omitempty" yaml:"rootVolumeEncryptionEnabled,omitempty"`

	// The tags for the WorkSpace. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The user name of the user for the WorkSpace. This user name must exist in the directory for the WorkSpace.
	UserName string `json:"userName,omitempty" yaml:"userName,omitempty"`
}
