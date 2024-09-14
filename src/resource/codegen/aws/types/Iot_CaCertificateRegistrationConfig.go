package types

type Iot_CaCertificateRegistrationConfig struct {
	// The ARN of the role.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The template body.
	TemplateBody string `json:"templateBody,omitempty" yaml:"templateBody,omitempty"`

	// The name of the provisioning template.
	TemplateName string `json:"templateName,omitempty" yaml:"templateName,omitempty"`
}
