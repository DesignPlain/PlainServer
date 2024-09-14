package alloydb

type User struct {
	// The database role name of the user.
	UserId string `json:"userId,omitempty" yaml:"userId,omitempty"`

	/*
	   The type of this user.
	   Possible values are: `ALLOYDB_BUILT_IN`, `ALLOYDB_IAM_USER`.


	   - - -
	*/
	UserType string `json:"userType,omitempty" yaml:"userType,omitempty"`

	/*
	   Identifies the alloydb cluster. Must be in the format
	   'projects/{project}/locations/{location}/clusters/{cluster_id}'
	*/
	Cluster string `json:"cluster,omitempty" yaml:"cluster,omitempty"`

	// List of database roles this database user has.
	DatabaseRoles []string `json:"databaseRoles,omitempty" yaml:"databaseRoles,omitempty"`

	// Password for this database user.
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
}
