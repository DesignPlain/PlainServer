package types

type Cognito_UserPoolUserAttributeUpdateSettings struct {
	// A list of attributes requiring verification before update. If set, the provided value(s) must also be set in `auto_verified_attributes`. Valid values: `email`, `phone_number`.
	AttributesRequireVerificationBeforeUpdates []string `json:"attributesRequireVerificationBeforeUpdates,omitempty" yaml:"attributesRequireVerificationBeforeUpdates,omitempty"`
}
