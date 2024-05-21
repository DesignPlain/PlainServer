package types

type Identityplatform_ConfigSignInPhoneNumber struct {
	// Whether phone number auth is enabled for the project or not.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// A map of <test phone number, fake code> that can be used for phone auth testing.
	TestPhoneNumbers map[string]string `json:"testPhoneNumbers,omitempty" yaml:"testPhoneNumbers,omitempty"`
}
