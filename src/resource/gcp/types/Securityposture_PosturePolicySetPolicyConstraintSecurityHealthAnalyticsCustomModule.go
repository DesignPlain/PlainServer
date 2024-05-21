package types

type Securityposture_PosturePolicySetPolicyConstraintSecurityHealthAnalyticsCustomModule struct {
	/*
	   Custom module details.
	   Structure is documented below.
	*/
	Config Securityposture_PosturePolicySetPolicyConstraintSecurityHealthAnalyticsCustomModuleConfig `json:"config,omitempty" yaml:"config,omitempty"`

	/*
	   The display name of the Security Health Analytics custom module. This
	   display name becomes the finding category for all findings that are
	   returned by this custom module.
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   (Output)
	   A server generated id of custom module.
	*/
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	/*
	   The state of enablement for the module at its level of the resource hierarchy.
	   Possible values are: `ENABLEMENT_STATE_UNSPECIFIED`, `ENABLED`, `DISABLED`.
	*/
	ModuleEnablementState string `json:"moduleEnablementState,omitempty" yaml:"moduleEnablementState,omitempty"`
}
