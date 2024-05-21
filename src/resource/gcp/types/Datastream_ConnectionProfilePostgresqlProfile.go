package types

type Datastream_ConnectionProfilePostgresqlProfile struct {
	// Database for the PostgreSQL connection.
	Database string `json:"database,omitempty" yaml:"database,omitempty"`

	// Hostname for the PostgreSQL connection.
	Hostname string `json:"hostname,omitempty" yaml:"hostname,omitempty"`

	/*
	   Password for the PostgreSQL connection.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	// Port for the PostgreSQL connection.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// Username for the PostgreSQL connection.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}
