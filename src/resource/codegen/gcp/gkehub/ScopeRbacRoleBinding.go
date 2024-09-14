package gkehub

import types "libds/gcp/types"

type ScopeRbacRoleBinding struct {
	/*
	   Principal that is be authorized in the cluster (at least of one the oneof
	   is required). Updating one will unset the other automatically.
	   group is the group, as seen by the kubernetes cluster.
	*/
	Group string `json:"group,omitempty" yaml:"group,omitempty"`

	/*
	   Labels for this ScopeRBACRoleBinding.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Role to bind to the principal.
	   Structure is documented below.
	*/
	Role types.Gkehub_ScopeRbacRoleBindingRole `json:"role,omitempty" yaml:"role,omitempty"`

	// Id of the scope
	ScopeId string `json:"scopeId,omitempty" yaml:"scopeId,omitempty"`

	// The client-provided identifier of the RBAC Role Binding.
	ScopeRbacRoleBindingId string `json:"scopeRbacRoleBindingId,omitempty" yaml:"scopeRbacRoleBindingId,omitempty"`

	/*
	   Principal that is be authorized in the cluster (at least of one the oneof
	   is required). Updating one will unset the other automatically.
	   user is the name of the user as seen by the kubernetes cluster, example
	   "alice" or "alice@domain.tld"
	*/
	User string `json:"user,omitempty" yaml:"user,omitempty"`
}
