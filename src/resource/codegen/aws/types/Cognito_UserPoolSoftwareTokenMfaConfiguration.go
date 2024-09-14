package types

type Cognito_UserPoolSoftwareTokenMfaConfiguration struct {
	// Boolean whether to enable software token Multi-Factor (MFA) tokens, such as Time-based One-Time Password (TOTP). To disable software token MFA When `sms_configuration` is not present, the `mfa_configuration` argument must be set to `OFF` and the `software_token_mfa_configuration` configuration block must be fully removed.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
