package cognito

type UserPoolDomain struct {
	// The user pool ID.
	UserPoolId string `json:"userPoolId,omitempty" yaml:"userPoolId,omitempty"`

	// The ARN of an ISSUED ACM certificate in us-east-1 for a custom domain.
	CertificateArn string `json:"certificateArn,omitempty" yaml:"certificateArn,omitempty"`

	// For custom domains, this is the fully-qualified domain name, such as auth.example.com. For Amazon Cognito prefix domains, this is the prefix alone, such as auth.
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`
}
