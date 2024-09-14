package types

type Bedrock_AgentAgentActionGroupFunctionSchemaMemberFunctionsFunction struct {
	// Description of the function and its purpose.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Name for the function.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Parameters that the agent elicits from the user to fulfill the function. See `parameters` Block for details.
	Parameters []Bedrock_AgentAgentActionGroupFunctionSchemaMemberFunctionsFunctionParameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}
