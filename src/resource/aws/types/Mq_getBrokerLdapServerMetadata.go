package types

type Mq_getBrokerLdapServerMetadata struct {
	//
	RoleSearchSubtree bool `json:"roleSearchSubtree,omitempty" yaml:"roleSearchSubtree,omitempty"`

	//
	ServiceAccountPassword string `json:"serviceAccountPassword,omitempty" yaml:"serviceAccountPassword,omitempty"`

	//
	UserBase string `json:"userBase,omitempty" yaml:"userBase,omitempty"`

	//
	UserRoleName string `json:"userRoleName,omitempty" yaml:"userRoleName,omitempty"`

	//
	UserSearchMatching string `json:"userSearchMatching,omitempty" yaml:"userSearchMatching,omitempty"`

	//
	RoleName string `json:"roleName,omitempty" yaml:"roleName,omitempty"`

	//
	RoleBase string `json:"roleBase,omitempty" yaml:"roleBase,omitempty"`

	//
	RoleSearchMatching string `json:"roleSearchMatching,omitempty" yaml:"roleSearchMatching,omitempty"`

	//
	ServiceAccountUsername string `json:"serviceAccountUsername,omitempty" yaml:"serviceAccountUsername,omitempty"`

	//
	UserSearchSubtree bool `json:"userSearchSubtree,omitempty" yaml:"userSearchSubtree,omitempty"`

	//
	Hosts []string `json:"hosts,omitempty" yaml:"hosts,omitempty"`
}
