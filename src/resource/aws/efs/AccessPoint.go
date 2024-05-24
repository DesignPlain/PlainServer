package efs

import types "DesignSphere_Server/src/resource/aws/types"

type AccessPoint struct {
	// Directory on the Amazon EFS file system that the access point provides access to. Detailed below.
	RootDirectory types.Efs_AccessPointRootDirectory `json:"rootDirectory,omitempty" yaml:"rootDirectory,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// ID of the file system for which the access point is intended.
	FileSystemId string `json:"fileSystemId,omitempty" yaml:"fileSystemId,omitempty"`

	// Operating system user and group applied to all file system requests made using the access point. Detailed below.
	PosixUser types.Efs_AccessPointPosixUser `json:"posixUser,omitempty" yaml:"posixUser,omitempty"`
}
