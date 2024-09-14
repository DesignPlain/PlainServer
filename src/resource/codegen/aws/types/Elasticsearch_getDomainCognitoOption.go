package types

type Elasticsearch_getDomainCognitoOption struct {
	// Whether node to node encryption is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The Cognito Identity pool used by the domain.
	IdentityPoolId string `json:"identityPoolId,omitempty" yaml:"identityPoolId,omitempty"`

	// The IAM Role with the AmazonESCognitoAccess policy attached.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The Cognito User pool used by the domain.
	UserPoolId string `json:"userPoolId,omitempty" yaml:"userPoolId,omitempty"`
}
