package types

type Securityposture_PosturePolicySetPolicyConstraintSecurityHealthAnalyticsModule struct {
	// The name of the module eg: BIGQUERY_TABLE_CMEK_DISABLED.
	ModuleName string `json:"moduleName,omitempty" yaml:"moduleName,omitempty"`

	/*
	   The state of enablement for the module at its level of the resource hierarchy.
	   Possible values are: `ENABLEMENT_STATE_UNSPECIFIED`, `ENABLED`, `DISABLED`.
	*/
	ModuleEnablementState string `json:"moduleEnablementState,omitempty" yaml:"moduleEnablementState,omitempty"`
}
