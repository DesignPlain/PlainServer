package types

type Cognito_getUserPoolDeviceConfiguration struct {
	// - Whether a challenge is required on new devices.
	ChallengeRequiredOnNewDevice bool `json:"challengeRequiredOnNewDevice,omitempty" yaml:"challengeRequiredOnNewDevice,omitempty"`

	// - Whether devices are only remembered if the user prompts it.
	DeviceOnlyRememberedOnUserPrompt bool `json:"deviceOnlyRememberedOnUserPrompt,omitempty" yaml:"deviceOnlyRememberedOnUserPrompt,omitempty"`
}
