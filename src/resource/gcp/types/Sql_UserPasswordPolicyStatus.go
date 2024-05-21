package types

type Sql_UserPasswordPolicyStatus struct {
	// If true, user does not have login privileges.
	Locked bool `json:"locked,omitempty" yaml:"locked,omitempty"`

	// Password expiration duration with one week grace period.
	PasswordExpirationTime string `json:"passwordExpirationTime,omitempty" yaml:"passwordExpirationTime,omitempty"`
}
