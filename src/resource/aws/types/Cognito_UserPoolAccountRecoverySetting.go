package types

type Cognito_UserPoolAccountRecoverySetting struct {
	// List of Account Recovery Options of the following structure:
	RecoveryMechanisms []Cognito_UserPoolAccountRecoverySettingRecoveryMechanism `json:"recoveryMechanisms,omitempty" yaml:"recoveryMechanisms,omitempty"`
}
