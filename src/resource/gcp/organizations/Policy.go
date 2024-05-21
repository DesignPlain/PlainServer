package organizations

import types "DesignSphere_Server/src/resource/gcp/types"

type Policy struct {
	// A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.
	ListPolicy types.Organizations_PolicyListPolicy `json:"listPolicy,omitempty" yaml:"listPolicy,omitempty"`

	// The numeric ID of the organization to set the policy for.
	OrgId string `json:"orgId,omitempty" yaml:"orgId,omitempty"`

	/*
	   A restore policy is a constraint to restore the default policy. Structure is documented below.

	   > --Note:-- If none of [`boolean_policy`, `list_policy`, `restore_policy`] are defined the policy for a given constraint will
	   effectively be unset. This is represented in the UI as the constraint being 'Inherited'.

	   - - -
	*/
	RestorePolicy types.Organizations_PolicyRestorePolicy `json:"restorePolicy,omitempty" yaml:"restorePolicy,omitempty"`

	// Version of the Policy. Default version is 0.
	Version int `json:"version,omitempty" yaml:"version,omitempty"`

	/*
	   A boolean policy is a constraint that is either enforced or not. Structure is documented
	   below.
	*/
	BooleanPolicy types.Organizations_PolicyBooleanPolicy `json:"booleanPolicy,omitempty" yaml:"booleanPolicy,omitempty"`

	/*
	   The name of the Constraint the Policy is configuring, for example, `serviceuser.services`. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints).

	   - - -
	*/
	Constraint string `json:"constraint,omitempty" yaml:"constraint,omitempty"`
}
