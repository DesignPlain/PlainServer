package cloudformation

type Stack struct {
	// A map of Parameter structures that specify input parameters for the stack.
	Parameters map[string]string `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// Map of resource tags to associate with this stack. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Structure containing the template body (max size: 51,200 bytes).
	TemplateBody string `json:"templateBody,omitempty" yaml:"templateBody,omitempty"`

	// The amount of time that can pass before the stack status becomes `CREATE_FAILED`.
	TimeoutInMinutes int `json:"timeoutInMinutes,omitempty" yaml:"timeoutInMinutes,omitempty"`

	/*
	   A list of capabilities.
	   Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, or `CAPABILITY_AUTO_EXPAND`
	*/
	Capabilities []string `json:"capabilities,omitempty" yaml:"capabilities,omitempty"`

	// Stack name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Action to be taken if stack creation fails. This must be
	   one of: `DO_NOTHING`, `ROLLBACK`, or `DELETE`. Conflicts with `disable_rollback`.
	*/
	OnFailure string `json:"onFailure,omitempty" yaml:"onFailure,omitempty"`

	/*
	   Structure containing the stack policy body.
	   Conflicts w/ `policy_url`.
	*/
	PolicyBody string `json:"policyBody,omitempty" yaml:"policyBody,omitempty"`

	/*
	   Location of a file containing the stack policy.
	   Conflicts w/ `policy_body`.
	*/
	PolicyUrl string `json:"policyUrl,omitempty" yaml:"policyUrl,omitempty"`

	// Location of a file containing the template body (max size: 460,800 bytes).
	TemplateUrl string `json:"templateUrl,omitempty" yaml:"templateUrl,omitempty"`

	/*
	   Set to true to disable rollback of the stack if stack creation failed.
	   Conflicts with `on_failure`.
	*/
	DisableRollback bool `json:"disableRollback,omitempty" yaml:"disableRollback,omitempty"`

	// The ARN of an IAM role that AWS CloudFormation assumes to create the stack. If you don't specify a value, AWS CloudFormation uses the role that was previously associated with the stack. If no role is available, AWS CloudFormation uses a temporary session that is generated from your user credentials.
	IamRoleArn string `json:"iamRoleArn,omitempty" yaml:"iamRoleArn,omitempty"`

	// A list of SNS topic ARNs to publish stack related events.
	NotificationArns []string `json:"notificationArns,omitempty" yaml:"notificationArns,omitempty"`
}
