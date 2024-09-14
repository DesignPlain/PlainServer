package quicksight

import types "libds/aws/types"

type Folder struct {
	// Identifier for the folder.
	FolderId string `json:"folderId,omitempty" yaml:"folderId,omitempty"`

	// The type of folder. By default, it is `SHARED`. Valid values are: `SHARED`.
	FolderType string `json:"folderType,omitempty" yaml:"folderType,omitempty"`

	/*
	   Display name for the folder.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The Amazon Resource Name (ARN) for the parent folder. If not set, creates a root-level folder.
	ParentFolderArn string `json:"parentFolderArn,omitempty" yaml:"parentFolderArn,omitempty"`

	// A set of resource permissions on the folder. Maximum of 64 items. See permissions.
	Permissions []types.Quicksight_FolderPermission `json:"permissions,omitempty" yaml:"permissions,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// AWS account ID.
	AwsAccountId string `json:"awsAccountId,omitempty" yaml:"awsAccountId,omitempty"`
}
