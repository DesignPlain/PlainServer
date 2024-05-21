package types

type Databasemigrationservice_ConnectionProfileOracleForwardSshConnectivity struct {
	// Required. Hostname for the SSH tunnel.
	Hostname string `json:"hostname,omitempty" yaml:"hostname,omitempty"`

	/*
	   Input only. SSH password. Only one of `password` and `private_key` can be configured.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	// Port for the SSH tunnel, default value is 22.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	/*
	   Input only. SSH private key. Only one of `password` and `private_key` can be configured.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	PrivateKey string `json:"privateKey,omitempty" yaml:"privateKey,omitempty"`

	// Required. Username for the SSH tunnel.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}
