package types

type Opensearch_DomainCognitoOptions struct {
	// ID of the Cognito User Pool to use.
	UserPoolId string `json:"userPoolId,omitempty" yaml:"userPoolId,omitempty"`

	// Whether Amazon Cognito authentication with Dashboard is enabled or not. Default is `false`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// ID of the Cognito Identity Pool to use.
	IdentityPoolId string `json:"identityPoolId,omitempty" yaml:"identityPoolId,omitempty"`

	// ARN of the IAM role that has the AmazonOpenSearchServiceCognitoAccess policy attached.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`
}
