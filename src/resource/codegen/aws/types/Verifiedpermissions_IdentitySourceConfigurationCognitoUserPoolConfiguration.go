package types

type Verifiedpermissions_IdentitySourceConfigurationCognitoUserPoolConfiguration struct {
	// The type of entity that a policy store maps to groups from an Amazon Cognito user pool identity source. See Group Configuration below.
	GroupConfiguration Verifiedpermissions_IdentitySourceConfigurationCognitoUserPoolConfigurationGroupConfiguration `json:"groupConfiguration,omitempty" yaml:"groupConfiguration,omitempty"`

	// The Amazon Resource Name (ARN) of the Amazon Cognito user pool that contains the identities to be authorized.
	UserPoolArn string `json:"userPoolArn,omitempty" yaml:"userPoolArn,omitempty"`

	// The unique application client IDs that are associated with the specified Amazon Cognito user pool.
	ClientIds []string `json:"clientIds,omitempty" yaml:"clientIds,omitempty"`
}
