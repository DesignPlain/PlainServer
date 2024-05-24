package types

type Ssmincidents_ResponsePlanActionSsmAutomation struct {
	// The key-value pair to resolve dynamic parameter values when processing a Systems Manager Automation runbook.
	DynamicParameters map[string]string `json:"dynamicParameters,omitempty" yaml:"dynamicParameters,omitempty"`

	// The key-value pair parameters to use when the automation document runs. The following values are supported:
	Parameters []Ssmincidents_ResponsePlanActionSsmAutomationParameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// The Amazon Resource Name (ARN) of the role that the automation document assumes when it runs commands.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The account that the automation document runs in. This can be in either the management account or an application account.
	TargetAccount string `json:"targetAccount,omitempty" yaml:"targetAccount,omitempty"`

	// The automation document's name.
	DocumentName string `json:"documentName,omitempty" yaml:"documentName,omitempty"`

	// The version of the automation document to use at runtime.
	DocumentVersion string `json:"documentVersion,omitempty" yaml:"documentVersion,omitempty"`
}
