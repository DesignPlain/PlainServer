package apigee

type Flowhook struct {
	// Where in the API call flow the flow hook is invoked. Must be one of PreProxyFlowHook, PostProxyFlowHook, PreTargetFlowHook, or PostTargetFlowHook.
	FlowHookPoint string `json:"flowHookPoint,omitempty" yaml:"flowHookPoint,omitempty"`

	// The Apigee Organization associated with the environment
	OrgId string `json:"orgId,omitempty" yaml:"orgId,omitempty"`

	// Id of the Sharedflow attaching to a flowhook point.
	Sharedflow string `json:"sharedflow,omitempty" yaml:"sharedflow,omitempty"`

	// Flag that specifies whether execution should continue if the flow hook throws an exception. Set to true to continue execution. Set to false to stop execution if the flow hook throws an exception. Defaults to true.
	ContinueOnError bool `json:"continueOnError,omitempty" yaml:"continueOnError,omitempty"`

	// Description of the flow hook.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The resource ID of the environment.
	Environment string `json:"environment,omitempty" yaml:"environment,omitempty"`
}
