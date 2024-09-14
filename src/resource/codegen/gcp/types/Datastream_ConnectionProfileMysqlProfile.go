package types

type Datastream_ConnectionProfileMysqlProfile struct {
	// Hostname for the MySQL connection.
	Hostname string `json:"hostname,omitempty" yaml:"hostname,omitempty"`

	/*
	   Password for the MySQL connection.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	// Port for the MySQL connection.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	/*
	   SSL configuration for the MySQL connection.
	   Structure is documented below.
	*/
	SslConfig Datastream_ConnectionProfileMysqlProfileSslConfig `json:"sslConfig,omitempty" yaml:"sslConfig,omitempty"`

	// Username for the MySQL connection.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}
