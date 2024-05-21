package types

type Datastream_ConnectionProfileForwardSshConnectivity struct {
	// Hostname for the SSH tunnel.
	Hostname string `json:"hostname,omitempty" yaml:"hostname,omitempty"`

	/*
	   SSH password.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	// Port for the SSH tunnel.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	/*
	   SSH private key.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	PrivateKey string `json:"privateKey,omitempty" yaml:"privateKey,omitempty"`

	// Username for the SSH tunnel.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}
