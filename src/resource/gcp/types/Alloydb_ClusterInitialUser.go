package types

type Alloydb_ClusterInitialUser struct {
	/*
	   The initial password for the user.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	// The database username.
	User string `json:"user,omitempty" yaml:"user,omitempty"`
}
