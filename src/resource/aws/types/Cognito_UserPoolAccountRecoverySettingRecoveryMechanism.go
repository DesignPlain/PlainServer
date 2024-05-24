package types

type Cognito_UserPoolAccountRecoverySettingRecoveryMechanism struct {
	/*
	   Name of the user pool.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Positive integer specifying priority of a method with 1 being the highest priority.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`
}
