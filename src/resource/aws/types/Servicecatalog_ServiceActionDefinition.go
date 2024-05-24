package types

type Servicecatalog_ServiceActionDefinition struct {
	// ARN of the role that performs the self-service actions on your behalf. For example, `arn:aws:iam::12345678910:role/ActionRole`. To reuse the provisioned product launch role, set to `LAUNCH_ROLE`.
	AssumeRole string `json:"assumeRole,omitempty" yaml:"assumeRole,omitempty"`

	// Name of the SSM document. For example, `AWS-RestartEC2Instance`. If you are using a shared SSM document, you must provide the ARN instead of the name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// List of parameters in JSON format. For example: `[{\"Name\":\"InstanceId\",\"Type\":\"TARGET\"}]` or `[{\"Name\":\"InstanceId\",\"Type\":\"TEXT_VALUE\"}]`.
	Parameters string `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// Service action definition type. Valid value is `SSM_AUTOMATION`. Default is `SSM_AUTOMATION`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// SSM document version. For example, `1`.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
