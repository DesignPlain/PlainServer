package redshift

type AuthenticationProfile struct {
	// The content of the authentication profile in JSON format. The maximum length of the JSON string is determined by a quota for your account.
	AuthenticationProfileContent string `json:"authenticationProfileContent,omitempty" yaml:"authenticationProfileContent,omitempty"`

	// The name of the authentication profile.
	AuthenticationProfileName string `json:"authenticationProfileName,omitempty" yaml:"authenticationProfileName,omitempty"`
}
