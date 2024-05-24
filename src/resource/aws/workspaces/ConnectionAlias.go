package workspaces

import types "DesignSphere_Server/src/resource/aws/types"

type ConnectionAlias struct {
	// The connection string specified for the connection alias. The connection string must be in the form of a fully qualified domain name (FQDN), such as www.example.com.
	ConnectionString string `json:"connectionString,omitempty" yaml:"connectionString,omitempty"`

	// A map of tags assigned to the WorkSpaces Connection Alias. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	//
	Timeouts types.Workspaces_ConnectionAliasTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`
}
