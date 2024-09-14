package gkehub

import types "libds/gcp/types"

type MembershipRbacRoleBinding struct {
	// Location of the Membership
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Id of the membership
	MembershipId string `json:"membershipId,omitempty" yaml:"membershipId,omitempty"`

	// The client-provided identifier of the RBAC Role Binding.
	MembershipRbacRoleBindingId string `json:"membershipRbacRoleBindingId,omitempty" yaml:"membershipRbacRoleBindingId,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Role to bind to the principal.
	   Structure is documented below.
	*/
	Role types.Gkehub_MembershipRbacRoleBindingRole `json:"role,omitempty" yaml:"role,omitempty"`

	/*
	   Principal that is be authorized in the cluster (at least of one the oneof
	   is required). Updating one will unset the other automatically.
	   user is the name of the user as seen by the kubernetes cluster, example
	   "alice" or "alice@domain.tld"
	*/
	User string `json:"user,omitempty" yaml:"user,omitempty"`
}
