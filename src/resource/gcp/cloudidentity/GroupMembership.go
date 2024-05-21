package cloudidentity

import types "DesignSphere_Server/src/resource/gcp/types"

type GroupMembership struct {
	/*
	   EntityKey of the member.
	   Structure is documented below.
	*/
	PreferredMemberKey types.Cloudidentity_GroupMembershipPreferredMemberKey `json:"preferredMemberKey,omitempty" yaml:"preferredMemberKey,omitempty"`

	/*
	   The MembershipRoles that apply to the Membership.
	   Must not contain duplicate MembershipRoles with the same name.
	   Structure is documented below.
	*/
	Roles []types.Cloudidentity_GroupMembershipRole `json:"roles,omitempty" yaml:"roles,omitempty"`

	// The name of the Group to create this membership in.
	Group string `json:"group,omitempty" yaml:"group,omitempty"`

	/*
	   EntityKey of the member.
	   Structure is documented below.
	*/
	MemberKey types.Cloudidentity_GroupMembershipMemberKey `json:"memberKey,omitempty" yaml:"memberKey,omitempty"`
}
