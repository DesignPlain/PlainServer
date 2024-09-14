package types

type Cognito_UserPoolPasswordPolicy struct {
	/*
	   Number of previous passwords that you want Amazon Cognito to restrict each user from reusing. Users can't set a password that matches any of number of previous passwords specified by this argument. A value of 0 means that password history is not enforced. Valid values are between 0 and 24.

	   --Note:-- This argument requires advanced security features to be active in the user pool.
	*/
	PasswordHistorySize int `json:"passwordHistorySize,omitempty" yaml:"passwordHistorySize,omitempty"`

	// Whether you have required users to use at least one lowercase letter in their password.
	RequireLowercase bool `json:"requireLowercase,omitempty" yaml:"requireLowercase,omitempty"`

	// Whether you have required users to use at least one number in their password.
	RequireNumbers bool `json:"requireNumbers,omitempty" yaml:"requireNumbers,omitempty"`

	// Whether you have required users to use at least one symbol in their password.
	RequireSymbols bool `json:"requireSymbols,omitempty" yaml:"requireSymbols,omitempty"`

	// Whether you have required users to use at least one uppercase letter in their password.
	RequireUppercase bool `json:"requireUppercase,omitempty" yaml:"requireUppercase,omitempty"`

	// In the password policy you have set, refers to the number of days a temporary password is valid. If the user does not sign-in during this time, their password will need to be reset by an administrator.
	TemporaryPasswordValidityDays int `json:"temporaryPasswordValidityDays,omitempty" yaml:"temporaryPasswordValidityDays,omitempty"`

	// Minimum length of the password policy that you have set.
	MinimumLength int `json:"minimumLength,omitempty" yaml:"minimumLength,omitempty"`
}
