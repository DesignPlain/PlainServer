package directoryservice

import types "DesignSphere_Server/src/resource/aws/types"

type SharedDirectory struct {
	/*
	   Identifier for the directory consumer account with whom the directory is to be shared. See below.

	   The following arguments are optional:
	*/
	Target types.Directoryservice_SharedDirectoryTarget `json:"target,omitempty" yaml:"target,omitempty"`

	// Identifier of the Managed Microsoft AD directory that you want to share with other accounts.
	DirectoryId string `json:"directoryId,omitempty" yaml:"directoryId,omitempty"`

	// Method used when sharing a directory. Valid values are `ORGANIZATIONS` and `HANDSHAKE`. Default is `HANDSHAKE`.
	Method string `json:"method,omitempty" yaml:"method,omitempty"`

	// Message sent by the directory owner to the directory consumer to help the directory consumer administrator determine whether to approve or reject the share invitation.
	Notes string `json:"notes,omitempty" yaml:"notes,omitempty"`
}
