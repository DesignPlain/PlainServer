package opensearch

import types "DesignSphere_Server/src/resource/aws/types"

type ServerlessSecurityConfig struct {
	// Name of the policy.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Configuration block for SAML options.
	SamlOptions types.Opensearch_ServerlessSecurityConfigSamlOptions `json:"samlOptions,omitempty" yaml:"samlOptions,omitempty"`

	/*
	   Type of configuration. Must be `saml`.

	   The following arguments are optional:
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Description of the security configuration.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
