package identitystore

type GroupMembership struct {
	// Identity Store ID associated with the Single Sign-On Instance.
	IdentityStoreId string `json:"identityStoreId,omitempty" yaml:"identityStoreId,omitempty"`

	// The identifier for a user in the Identity Store.
	MemberId string `json:"memberId,omitempty" yaml:"memberId,omitempty"`

	// The identifier for a group in the Identity Store.
	GroupId string `json:"groupId,omitempty" yaml:"groupId,omitempty"`
}
