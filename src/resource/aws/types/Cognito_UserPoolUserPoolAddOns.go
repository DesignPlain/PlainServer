package types

type Cognito_UserPoolUserPoolAddOns struct {
	// Mode for advanced security, must be one of `OFF`, `AUDIT` or `ENFORCED`.
	AdvancedSecurityMode string `json:"advancedSecurityMode,omitempty" yaml:"advancedSecurityMode,omitempty"`
}
