package types

type Cognito_getUserPoolAccountRecoverySettingRecoveryMechanism struct {
	// - Name of the attribute.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// - Priority of this mechanism in the recovery process (lower numbers are higher priority).
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`
}
