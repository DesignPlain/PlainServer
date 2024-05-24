package types

type Cognito_UserPoolUsernameConfiguration struct {
	// Whether username case sensitivity will be applied for all users in the user pool through Cognito APIs.
	CaseSensitive bool `json:"caseSensitive,omitempty" yaml:"caseSensitive,omitempty"`
}
