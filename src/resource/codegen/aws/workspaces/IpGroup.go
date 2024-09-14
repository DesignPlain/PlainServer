package workspaces

import types "libds/aws/types"

type IpGroup struct {
	// One or more pairs specifying the IP group rule (in CIDR format) from which web requests originate.
	Rules []types.Workspaces_IpGroupRule `json:"rules,omitempty" yaml:"rules,omitempty"`

	// A map of tags assigned to the WorkSpaces directory. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The description of the IP group.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the IP group.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
