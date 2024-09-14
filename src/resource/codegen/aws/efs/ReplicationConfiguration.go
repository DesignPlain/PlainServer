package efs

import types "libds/aws/types"

type ReplicationConfiguration struct {
	// A destination configuration block (documented below).
	Destination types.Efs_ReplicationConfigurationDestination `json:"destination,omitempty" yaml:"destination,omitempty"`

	// The ID of the file system that is to be replicated.
	SourceFileSystemId string `json:"sourceFileSystemId,omitempty" yaml:"sourceFileSystemId,omitempty"`
}
