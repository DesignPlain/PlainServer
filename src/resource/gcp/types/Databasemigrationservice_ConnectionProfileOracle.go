package types

type Databasemigrationservice_ConnectionProfileOracle struct {
	// Required. The IP or hostname of the source Oracle database.
	Host string `json:"host,omitempty" yaml:"host,omitempty"`

	/*
	   Configuration for using a private network to communicate with the source database
	   Structure is documented below.
	*/
	PrivateConnectivity Databasemigrationservice_ConnectionProfileOraclePrivateConnectivity `json:"privateConnectivity,omitempty" yaml:"privateConnectivity,omitempty"`

	// Required. The username that Database Migration Service will use to connect to the database. The value is encrypted when stored in Database Migration Service.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`

	// Required. Database service for the Oracle connection.
	DatabaseService string `json:"databaseService,omitempty" yaml:"databaseService,omitempty"`

	/*
	   SSL configuration for the destination to connect to the source database.
	   Structure is documented below.
	*/
	ForwardSshConnectivity Databasemigrationservice_ConnectionProfileOracleForwardSshConnectivity `json:"forwardSshConnectivity,omitempty" yaml:"forwardSshConnectivity,omitempty"`

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

	// Required. The network port of the source Oracle database.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	/*
	   SSL configuration for the destination to connect to the source database.
	   Structure is documented below.
	*/
	Ssl Databasemigrationservice_ConnectionProfileOracleSsl `json:"ssl,omitempty" yaml:"ssl,omitempty"`

	/*
	   This object has no nested fields.
	   Static IP address connectivity configured on service project.
	*/
	StaticServiceIpConnectivity Databasemigrationservice_ConnectionProfileOracleStaticServiceIpConnectivity `json:"staticServiceIpConnectivity,omitempty" yaml:"staticServiceIpConnectivity,omitempty"`
}
