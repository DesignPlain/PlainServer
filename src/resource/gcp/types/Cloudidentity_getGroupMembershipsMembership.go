package types

type Cloudidentity_getGroupMembershipsMembership struct {
	// The MembershipRoles that apply to the Membership. Structure is documented below.
	Roles []Cloudidentity_getGroupMembershipsMembershipRole `json:"roles,omitempty" yaml:"roles,omitempty"`

	// The type of the membership.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The time when the Membership was last updated.
	UpdateTime string `json:"updateTime,omitempty" yaml:"updateTime,omitempty"`

	// The time when the Membership was created.
	CreateTime string `json:"createTime,omitempty" yaml:"createTime,omitempty"`

	// The parent Group resource under which to lookup the Membership names. Must be of the form groups/{group_id}.
	Group string `json:"group,omitempty" yaml:"group,omitempty"`

	// EntityKey of the member.  Structure is documented below.
	MemberKeys []Cloudidentity_getGroupMembershipsMembershipMemberKey `json:"memberKeys,omitempty" yaml:"memberKeys,omitempty"`

	// The name of the MembershipRole. One of OWNER, MANAGER, MEMBER.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// EntityKey of the member.  Structure is documented below.
	PreferredMemberKeys []Cloudidentity_getGroupMembershipsMembershipPreferredMemberKey `json:"preferredMemberKeys,omitempty" yaml:"preferredMemberKeys,omitempty"`
}
