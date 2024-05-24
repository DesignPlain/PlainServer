package types

type Cleanrooms_CollaborationMember struct {
	//
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	//
	MemberAbilities []string `json:"memberAbilities,omitempty" yaml:"memberAbilities,omitempty"`

	//
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	//
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`
}
