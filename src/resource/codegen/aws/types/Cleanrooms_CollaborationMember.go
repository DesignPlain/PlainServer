package types

type Cleanrooms_CollaborationMember struct {
	//
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	//
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	//
	MemberAbilities []string `json:"memberAbilities,omitempty" yaml:"memberAbilities,omitempty"`

	//
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
