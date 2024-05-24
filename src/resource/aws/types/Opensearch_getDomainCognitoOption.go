package types

type Opensearch_getDomainCognitoOption struct {
	// Enabled disabled toggle for off-peak update window
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Cognito Identity pool used by the domain.
	IdentityPoolId string `json:"identityPoolId,omitempty" yaml:"identityPoolId,omitempty"`

	// IAM Role with the AmazonOpenSearchServiceCognitoAccess policy attached.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// Cognito User pool used by the domain.
	UserPoolId string `json:"userPoolId,omitempty" yaml:"userPoolId,omitempty"`
}
