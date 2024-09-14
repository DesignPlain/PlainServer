package types

type Cognito_UserPoolAccountRecoverySettingRecoveryMechanism struct {
	// Recovery method for a user. Can be of the following: `verified_email`, `verified_phone_number`, and `admin_only`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Positive integer specifying priority of a method with 1 being the highest priority.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`
}
