package types

type Cloudidentity_GroupMembershipRole struct {
	/*
	   The MembershipRole expiry details, only supported for MEMBER role.
	   Other roles cannot be accompanied with MEMBER role having expiry.
	   Structure is documented below.
	*/
	ExpiryDetail Cloudidentity_GroupMembershipRoleExpiryDetail `json:"expiryDetail,omitempty" yaml:"expiryDetail,omitempty"`

	/*
	   The name of the MembershipRole. Must be one of OWNER, MANAGER, MEMBER.
	   Possible values are: `OWNER`, `MANAGER`, `MEMBER`.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
