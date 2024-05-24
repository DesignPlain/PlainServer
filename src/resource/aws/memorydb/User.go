package memorydb

import types "DesignSphere_Server/src/resource/aws/types"

type User struct {
	// The access permissions string used for this user.
	AccessString string `json:"accessString,omitempty" yaml:"accessString,omitempty"`

	// Denotes the user's authentication properties. Detailed below.
	AuthenticationMode types.Memorydb_UserAuthenticationMode `json:"authenticationMode,omitempty" yaml:"authenticationMode,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   Name of the MemoryDB user. Up to 40 characters.

	   The following arguments are optional:
	*/
	UserName string `json:"userName,omitempty" yaml:"userName,omitempty"`
}
