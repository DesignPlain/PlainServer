package types

type Cloudidentity_getGroupMembershipsMembershipRole struct {
	/*
	   The MembershipRole expiry details, only supported for MEMBER role.
	   Other roles cannot be accompanied with MEMBER role having expiry.
	*/
	ExpiryDetails []Cloudidentity_getGroupMembershipsMembershipRoleExpiryDetail `json:"expiryDetails,omitempty" yaml:"expiryDetails,omitempty"`

	// The name of the MembershipRole. One of OWNER, MANAGER, MEMBER.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
