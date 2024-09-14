package types

type Verifiedpermissions_IdentitySourceConfigurationOpenIdConnectConfigurationGroupConfiguration struct {
	// The token claim that you want Verified Permissions to interpret as group membership. For example, `groups`.
	GroupClaim string `json:"groupClaim,omitempty" yaml:"groupClaim,omitempty"`

	// The name of the schema entity type that's mapped to the user pool group. Defaults to `AWS::CognitoGroup`.
	GroupEntityType string `json:"groupEntityType,omitempty" yaml:"groupEntityType,omitempty"`
}
