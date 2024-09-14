package types

type Iam_AccessBoundaryPolicyRuleAccessBoundaryRule struct {
	/*
	   The availability condition further constrains the access allowed by the access boundary rule.
	   Structure is documented below.
	*/
	AvailabilityCondition Iam_AccessBoundaryPolicyRuleAccessBoundaryRuleAvailabilityCondition `json:"availabilityCondition,omitempty" yaml:"availabilityCondition,omitempty"`

	// A list of permissions that may be allowed for use on the specified resource.
	AvailablePermissions []string `json:"availablePermissions,omitempty" yaml:"availablePermissions,omitempty"`

	// The full resource name of a Google Cloud resource entity.
	AvailableResource string `json:"availableResource,omitempty" yaml:"availableResource,omitempty"`
}
