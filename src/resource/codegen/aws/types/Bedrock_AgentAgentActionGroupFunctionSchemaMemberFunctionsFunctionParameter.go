package types

type Bedrock_AgentAgentActionGroupFunctionSchemaMemberFunctionsFunctionParameter struct {
	// Description of the parameter. Helps the foundation model determine how to elicit the parameters from the user.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Name of the parameter.

	   --Note:-- The argument name `map_block_key` may seem out of context, but is necessary for backward compatibility reasons in the provider.
	*/
	MapBlockKey string `json:"mapBlockKey,omitempty" yaml:"mapBlockKey,omitempty"`

	// Whether the parameter is required for the agent to complete the function for action group invocation.
	Required bool `json:"required,omitempty" yaml:"required,omitempty"`

	// Data type of the parameter. Valid values: `string`, `number`, `integer`, `boolean`, `array`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
