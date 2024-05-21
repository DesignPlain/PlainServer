package types

type Sql_UserSqlServerUserDetail struct {
	// If the user has been disabled.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`

	// The server roles for this user in the database.
	ServerRoles []string `json:"serverRoles,omitempty" yaml:"serverRoles,omitempty"`
}
