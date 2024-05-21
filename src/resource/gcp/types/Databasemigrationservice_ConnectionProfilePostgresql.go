package types

type Databasemigrationservice_ConnectionProfilePostgresql struct {
	/*
	   (Output)
	   Output only. If the source is a Cloud SQL database, this field indicates the network architecture it's associated with.
	*/
	NetworkArchitecture string `json:"networkArchitecture,omitempty" yaml:"networkArchitecture,omitempty"`

	/*
	   Required. Input only. The password for the user that Database Migration Service will be using to connect to the database.
	   This field is not returned on request, and the value is encrypted when stored in Database Migration Service.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	/*
	   (Output)
	   Output only. Indicates If this connection profile password is stored.
	*/
	PasswordSet bool `json:"passwordSet,omitempty" yaml:"passwordSet,omitempty"`

	// Required. The network port of the source MySQL database.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	/*
	   SSL configuration for the destination to connect to the source database.
	   Structure is documented below.
	*/
	Ssl Databasemigrationservice_ConnectionProfilePostgresqlSsl `json:"ssl,omitempty" yaml:"ssl,omitempty"`

	// Required. The username that Database Migration Service will use to connect to the database. The value is encrypted when stored in Database Migration Service.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`

	// If the source is a Cloud SQL database, use this field to provide the Cloud SQL instance ID of the source.
	CloudSqlId string `json:"cloudSqlId,omitempty" yaml:"cloudSqlId,omitempty"`

	// Required. The IP or hostname of the source MySQL database.
	Host string `json:"host,omitempty" yaml:"host,omitempty"`
}
