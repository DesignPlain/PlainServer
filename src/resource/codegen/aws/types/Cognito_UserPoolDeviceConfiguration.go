package types

type Cognito_UserPoolDeviceConfiguration struct {
	// Whether a device is only remembered on user prompt. `false` equates to "Always" remember, `true` is "User Opt In," and not using a `device_configuration` block is "No."
	DeviceOnlyRememberedOnUserPrompt bool `json:"deviceOnlyRememberedOnUserPrompt,omitempty" yaml:"deviceOnlyRememberedOnUserPrompt,omitempty"`

	// Whether a challenge is required on a new device. Only applicable to a new device.
	ChallengeRequiredOnNewDevice bool `json:"challengeRequiredOnNewDevice,omitempty" yaml:"challengeRequiredOnNewDevice,omitempty"`
}
