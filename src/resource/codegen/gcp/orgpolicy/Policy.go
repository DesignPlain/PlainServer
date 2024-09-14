package orgpolicy

import types "libds/gcp/types"

type Policy struct {
	// Dry-run policy. Audit-only policy, can be used to monitor how the policy would have impacted the existing and future resources if it's enforced.
	DryRunSpec types.Orgpolicy_PolicyDryRunSpec `json:"dryRunSpec,omitempty" yaml:"dryRunSpec,omitempty"`

	// Immutable. The resource name of the Policy. Must be one of the following forms, where constraint_name is the name of the constraint which this Policy configures: - `projects/{project_number}/policies/{constraint_name}` - `folders/{folder_id}/policies/{constraint_name}` - `organizations/{organization_id}/policies/{constraint_name}` For example, "projects/123/policies/compute.disableSerialPortAccess". Note: `projects/{project_id}/policies/{constraint_name}` is also an acceptable name for API requests, but responses will return the name using the equivalent project number.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The parent of the resource.



	   - - -
	*/
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`

	// Basic information about the Organization Policy.
	Spec types.Orgpolicy_PolicySpec `json:"spec,omitempty" yaml:"spec,omitempty"`
}
