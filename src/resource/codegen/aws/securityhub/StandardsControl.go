package securityhub

type StandardsControl struct {
	// The control status could be `ENABLED` or `DISABLED`. You have to specify `disabled_reason` argument for `DISABLED` control status.
	ControlStatus string `json:"controlStatus,omitempty" yaml:"controlStatus,omitempty"`

	// A description of the reason why you are disabling a security standard control. If you specify this attribute, `control_status` will be set to `DISABLED` automatically.
	DisabledReason string `json:"disabledReason,omitempty" yaml:"disabledReason,omitempty"`

	// The standards control ARN. See the AWS documentation for how to list existing controls using [`get-enabled-standards`](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/securityhub/get-enabled-standards.html) and [`describe-standards-controls`](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/securityhub/describe-standards-controls.html).
	StandardsControlArn string `json:"standardsControlArn,omitempty" yaml:"standardsControlArn,omitempty"`
}
