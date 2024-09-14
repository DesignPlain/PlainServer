package medialive

import types "libds/aws/types"

type InputSecurityGroup struct {
	// A map of tags to assign to the InputSecurityGroup. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   Whitelist rules. See Whitelist Rules for more details.

	   The following arguments are optional:
	*/
	WhitelistRules []types.Medialive_InputSecurityGroupWhitelistRule `json:"whitelistRules,omitempty" yaml:"whitelistRules,omitempty"`
}
