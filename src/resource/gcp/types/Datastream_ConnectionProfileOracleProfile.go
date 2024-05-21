package types

type Datastream_ConnectionProfileOracleProfile struct {
	// Username for the Oracle connection.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`

	// Connection string attributes
	ConnectionAttributes map[string]string `json:"connectionAttributes,omitempty" yaml:"connectionAttributes,omitempty"`

	// Database for the Oracle connection.
	DatabaseService string `json:"databaseService,omitempty" yaml:"databaseService,omitempty"`

	// Hostname for the Oracle connection.
	Hostname string `json:"hostname,omitempty" yaml:"hostname,omitempty"`

	/*
	   Password for the Oracle connection.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	// Port for the Oracle connection.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
}
