package efs

import types "DesignSphere_Server/src/resource/aws/types"

type ReplicationConfiguration struct {
	// A destination configuration block (documented below).
	Destination types.Efs_ReplicationConfigurationDestination `json:"destination,omitempty" yaml:"destination,omitempty"`

	// The ID of the file system that is to be replicated.
	SourceFileSystemId string `json:"sourceFileSystemId,omitempty" yaml:"sourceFileSystemId,omitempty"`
}
