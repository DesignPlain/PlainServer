package sql

type SourceRepresentationInstance struct {
	// The CA certificate on the external server. Include only if SSL/TLS is used on the external server.
	CaCertificate string `json:"caCertificate,omitempty" yaml:"caCertificate,omitempty"`

	// The client certificate on the external server. Required only for server-client authentication. Include only if SSL/TLS is used on the external server.
	ClientCertificate string `json:"clientCertificate,omitempty" yaml:"clientCertificate,omitempty"`

	/*
	   The IPv4 address and port for the external server, or the the DNS address for the external server. If the external server is hosted on Cloud SQL, the port is 5432.


	   - - -
	*/
	Host string `json:"host,omitempty" yaml:"host,omitempty"`

	/*
	   The externally accessible port for the source database server.
	   Defaults to 3306.
	*/
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The private key file for the client certificate on the external server. Required only for server-client authentication. Include only if SSL/TLS is used on the external server.
	ClientKey string `json:"clientKey,omitempty" yaml:"clientKey,omitempty"`

	/*
	   The MySQL version running on your source database server.
	   Possible values are: `MYSQL_5_6`, `MYSQL_5_7`, `MYSQL_8_0`, `POSTGRES_9_6`, `POSTGRES_10`, `POSTGRES_11`, `POSTGRES_12`, `POSTGRES_13`, `POSTGRES_14`.
	*/
	DatabaseVersion string `json:"databaseVersion,omitempty" yaml:"databaseVersion,omitempty"`

	// A file in the bucket that contains the data from the external server.
	DumpFilePath string `json:"dumpFilePath,omitempty" yaml:"dumpFilePath,omitempty"`

	// The name of the source representation instance. Use any valid Cloud SQL instance name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The password for the replication user account.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	/*
	   The Region in which the created instance should reside.
	   If it is not provided, the provider region is used.
	*/
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// The replication user account on the external server.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}
