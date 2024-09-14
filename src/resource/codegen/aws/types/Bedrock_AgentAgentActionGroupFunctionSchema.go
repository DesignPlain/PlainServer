package types

type Bedrock_AgentAgentActionGroupFunctionSchema struct {
	/*
	   Contains a list of functions.
	   Each function describes and action in the action group.
	   See `member_functions` Block for details.
	*/
	MemberFunctions Bedrock_AgentAgentActionGroupFunctionSchemaMemberFunctions `json:"memberFunctions,omitempty" yaml:"memberFunctions,omitempty"`
}
