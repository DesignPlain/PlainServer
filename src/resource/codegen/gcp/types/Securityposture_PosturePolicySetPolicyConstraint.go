package types

type Securityposture_PosturePolicySetPolicyConstraint struct {
	/*
	   Definition of Security Health Analytics Custom Module.
	   Structure is documented below.
	*/
	SecurityHealthAnalyticsCustomModule Securityposture_PosturePolicySetPolicyConstraintSecurityHealthAnalyticsCustomModule `json:"securityHealthAnalyticsCustomModule,omitempty" yaml:"securityHealthAnalyticsCustomModule,omitempty"`

	/*
	   Security Health Analytics built-in detector definition.
	   Structure is documented below.
	*/
	SecurityHealthAnalyticsModule Securityposture_PosturePolicySetPolicyConstraintSecurityHealthAnalyticsModule `json:"securityHealthAnalyticsModule,omitempty" yaml:"securityHealthAnalyticsModule,omitempty"`

	/*
	   Organization policy canned constraint definition.
	   Structure is documented below.
	*/
	OrgPolicyConstraint Securityposture_PosturePolicySetPolicyConstraintOrgPolicyConstraint `json:"orgPolicyConstraint,omitempty" yaml:"orgPolicyConstraint,omitempty"`

	/*
	   Organization policy custom constraint policy definition.
	   Structure is documented below.
	*/
	OrgPolicyConstraintCustom Securityposture_PosturePolicySetPolicyConstraintOrgPolicyConstraintCustom `json:"orgPolicyConstraintCustom,omitempty" yaml:"orgPolicyConstraintCustom,omitempty"`
}
