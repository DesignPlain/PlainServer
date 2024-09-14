package types

type Apprunner_ServiceSourceConfigurationAuthenticationConfiguration struct {
	// ARN of the IAM role that grants the App Runner service access to a source repository. Required for ECR image repositories (but not for ECR Public)
	AccessRoleArn string `json:"accessRoleArn,omitempty" yaml:"accessRoleArn,omitempty"`

	// ARN of the App Runner connection that enables the App Runner service to connect to a source repository. Required for GitHub code repositories.
	ConnectionArn string `json:"connectionArn,omitempty" yaml:"connectionArn,omitempty"`
}
