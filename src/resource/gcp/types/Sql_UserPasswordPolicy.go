package types

type Sql_UserPasswordPolicy struct {
	// Password expiration duration with one week grace period.
	PasswordExpirationDuration string `json:"passwordExpirationDuration,omitempty" yaml:"passwordExpirationDuration,omitempty"`

	//
	Statuses []Sql_UserPasswordPolicyStatus `json:"statuses,omitempty" yaml:"statuses,omitempty"`

	// Number of failed attempts allowed before the user get locked.
	AllowedFailedAttempts int `json:"allowedFailedAttempts,omitempty" yaml:"allowedFailedAttempts,omitempty"`

	// If true, the check that will lock user after too many failed login attempts will be enabled.
	EnableFailedAttemptsCheck bool `json:"enableFailedAttemptsCheck,omitempty" yaml:"enableFailedAttemptsCheck,omitempty"`

	// If true, the user must specify the current password before changing the password. This flag is supported only for MySQL.
	EnablePasswordVerification bool `json:"enablePasswordVerification,omitempty" yaml:"enablePasswordVerification,omitempty"`
}
