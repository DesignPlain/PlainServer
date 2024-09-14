package types

type Verifiedpermissions_IdentitySourceConfigurationCognitoUserPoolConfigurationGroupConfiguration struct {
	// The name of the schema entity type that's mapped to the user pool group. Defaults to `AWS::CognitoGroup`.
	GroupEntityType string `json:"groupEntityType,omitempty" yaml:"groupEntityType,omitempty"`
}
