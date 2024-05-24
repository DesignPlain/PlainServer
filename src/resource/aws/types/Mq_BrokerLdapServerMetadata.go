package types

type Mq_BrokerLdapServerMetadata struct {
	// Search criteria for groups.
	RoleSearchMatching string `json:"roleSearchMatching,omitempty" yaml:"roleSearchMatching,omitempty"`

	// Service account password.
	ServiceAccountPassword string `json:"serviceAccountPassword,omitempty" yaml:"serviceAccountPassword,omitempty"`

	// Service account username.
	ServiceAccountUsername string `json:"serviceAccountUsername,omitempty" yaml:"serviceAccountUsername,omitempty"`

	// Fully qualified name of the directory where you want to search for users.
	UserBase string `json:"userBase,omitempty" yaml:"userBase,omitempty"`

	// Specifies the name of the LDAP attribute for the user group membership.
	UserRoleName string `json:"userRoleName,omitempty" yaml:"userRoleName,omitempty"`

	// List of a fully qualified domain name of the LDAP server and an optional failover server.
	Hosts []string `json:"hosts,omitempty" yaml:"hosts,omitempty"`

	// Specifies the LDAP attribute that identifies the group name attribute in the object returned from the group membership query.
	RoleName string `json:"roleName,omitempty" yaml:"roleName,omitempty"`

	// Whether the directory search scope is the entire sub-tree.
	RoleSearchSubtree bool `json:"roleSearchSubtree,omitempty" yaml:"roleSearchSubtree,omitempty"`

	// Search criteria for users.
	UserSearchMatching string `json:"userSearchMatching,omitempty" yaml:"userSearchMatching,omitempty"`

	// Whether the directory search scope is the entire sub-tree.
	UserSearchSubtree bool `json:"userSearchSubtree,omitempty" yaml:"userSearchSubtree,omitempty"`

	// Fully qualified name of the directory to search for a user’s groups.
	RoleBase string `json:"roleBase,omitempty" yaml:"roleBase,omitempty"`
}
