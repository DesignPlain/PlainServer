package types

type Mq_getBrokerLdapServerMetadata struct {
	//
	ServiceAccountUsername string `json:"serviceAccountUsername,omitempty" yaml:"serviceAccountUsername,omitempty"`

	//
	UserBase string `json:"userBase,omitempty" yaml:"userBase,omitempty"`

	//
	UserSearchMatching string `json:"userSearchMatching,omitempty" yaml:"userSearchMatching,omitempty"`

	//
	UserSearchSubtree bool `json:"userSearchSubtree,omitempty" yaml:"userSearchSubtree,omitempty"`

	//
	Hosts []string `json:"hosts,omitempty" yaml:"hosts,omitempty"`

	//
	RoleSearchMatching string `json:"roleSearchMatching,omitempty" yaml:"roleSearchMatching,omitempty"`

	//
	RoleSearchSubtree bool `json:"roleSearchSubtree,omitempty" yaml:"roleSearchSubtree,omitempty"`

	//
	ServiceAccountPassword string `json:"serviceAccountPassword,omitempty" yaml:"serviceAccountPassword,omitempty"`

	//
	RoleBase string `json:"roleBase,omitempty" yaml:"roleBase,omitempty"`

	//
	RoleName string `json:"roleName,omitempty" yaml:"roleName,omitempty"`

	//
	UserRoleName string `json:"userRoleName,omitempty" yaml:"userRoleName,omitempty"`
}
