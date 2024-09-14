package sfn

import types "libds/aws/types"

type Alias struct {
	// Description of the alias.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Name for the alias you are creating.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The StateMachine alias' route configuration settings. Fields documented below
	RoutingConfigurations []types.Sfn_AliasRoutingConfiguration `json:"routingConfigurations,omitempty" yaml:"routingConfigurations,omitempty"`
}
