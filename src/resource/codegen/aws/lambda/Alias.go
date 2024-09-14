package lambda

import types "libds/aws/types"

type Alias struct {
	// Name for the alias you are creating. Pattern: `(?!^[0-9]+$)([a-zA-Z0-9-_]+)`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The Lambda alias' route configuration settings. Fields documented below
	RoutingConfig types.Lambda_AliasRoutingConfig `json:"routingConfig,omitempty" yaml:"routingConfig,omitempty"`

	// Description of the alias.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Lambda Function name or ARN.
	FunctionName string `json:"functionName,omitempty" yaml:"functionName,omitempty"`

	// Lambda function version for which you are creating the alias. Pattern: `(\$LATEST|[0-9]+)`.
	FunctionVersion string `json:"functionVersion,omitempty" yaml:"functionVersion,omitempty"`
}
