package types

type Securityposture_PosturePolicySetPolicyConstraintOrgPolicyConstraintCustomCustomConstraint struct {
	// Immutable. The fully qualified name of the Google Cloud REST resource containing the object and field you want to restrict. For example, `container.googleapis.com/NodePool`.
	ResourceTypes []string `json:"resourceTypes,omitempty" yaml:"resourceTypes,omitempty"`

	/*
	   The action to take if the condition is met.
	   Possible values are: `ALLOW`, `DENY`.
	*/
	ActionType string `json:"actionType,omitempty" yaml:"actionType,omitempty"`

	// A CEL condition that refers to a supported service resource, for example `resource.management.autoUpgrade == false`. For details about CEL usage, see [Common Expression Language](https://cloud.google.com/resource-manager/docs/organization-policy/creating-managing-custom-constraints#common_expression_language).
	Condition string `json:"condition,omitempty" yaml:"condition,omitempty"`

	// A human-friendly description of the constraint to display as an error message when the policy is violated.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// A human-friendly name for the constraint.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// A list of RESTful methods for which to enforce the constraint. Can be `CREATE`, `UPDATE`, or both. Not all Google Cloud services support both methods. To see supported methods for each service, find the service in [Supported services](https://cloud.google.com/resource-manager/docs/organization-policy/custom-constraint-supported-services).
	MethodTypes []string `json:"methodTypes,omitempty" yaml:"methodTypes,omitempty"`

	// Immutable. The name of the custom constraint. This is unique within the organization.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
