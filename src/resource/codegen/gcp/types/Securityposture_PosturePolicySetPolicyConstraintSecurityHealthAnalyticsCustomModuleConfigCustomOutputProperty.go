package types

type Securityposture_PosturePolicySetPolicyConstraintSecurityHealthAnalyticsCustomModuleConfigCustomOutputProperty struct {
	// Name of the property for the custom output.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The CEL expression for the custom output. A resource property can be
	   specified to return the value of the property or a text string enclosed
	   in quotation marks.
	   Structure is documented below.
	*/
	ValueExpression Securityposture_PosturePolicySetPolicyConstraintSecurityHealthAnalyticsCustomModuleConfigCustomOutputPropertyValueExpression `json:"valueExpression,omitempty" yaml:"valueExpression,omitempty"`
}
