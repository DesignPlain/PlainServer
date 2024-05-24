package types

type Elasticsearch_DomainCognitoOptions struct {
	// Whether Amazon Cognito authentication with Kibana is enabled or not.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// ID of the Cognito Identity Pool to use.
	IdentityPoolId string `json:"identityPoolId,omitempty" yaml:"identityPoolId,omitempty"`

	// ARN of the IAM role that has the AmazonESCognitoAccess policy attached.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// ID of the Cognito User Pool to use.
	UserPoolId string `json:"userPoolId,omitempty" yaml:"userPoolId,omitempty"`
}
