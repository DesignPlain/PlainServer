package gamelift

import types "DesignSphere_Server/src/resource/aws/types"

type Alias struct {
	// Specifies the fleet and/or routing type to use for the alias.
	RoutingStrategy types.Gamelift_AliasRoutingStrategy `json:"routingStrategy,omitempty" yaml:"routingStrategy,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Description of the alias.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Name of the alias.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
