package types

type Sql_DatabaseInstanceReplicaConfiguration struct {
	// Username for replication connection.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`

	/*
	   True if the master's common name
	   value is checked during the SSL handshake.
	*/
	VerifyServerCertificate bool `json:"verifyServerCertificate,omitempty" yaml:"verifyServerCertificate,omitempty"`

	/*
	   Path to a SQL file in GCS from which replica
	   instances are created. Format is `gs://bucket/filename`.
	*/
	DumpFilePath string `json:"dumpFilePath,omitempty" yaml:"dumpFilePath,omitempty"`

	/*
	   Time in ms between replication
	   heartbeats.
	*/
	MasterHeartbeatPeriod int `json:"masterHeartbeatPeriod,omitempty" yaml:"masterHeartbeatPeriod,omitempty"`

	// Password for the replication connection.
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	// Permissible ciphers for use in SSL encryption.
	SslCipher string `json:"sslCipher,omitempty" yaml:"sslCipher,omitempty"`

	/*
	   Specifies if the replica is the failover target.
	   If the field is set to true the replica will be designated as a failover replica.
	   If the master instance fails, the replica instance will be promoted as
	   the new master instance.
	   > --NOTE:-- Not supported for Postgres database.
	*/
	FailoverTarget bool `json:"failoverTarget,omitempty" yaml:"failoverTarget,omitempty"`

	/*
	   PEM representation of the trusted CA's x509
	   certificate.
	*/
	CaCertificate string `json:"caCertificate,omitempty" yaml:"caCertificate,omitempty"`

	/*
	   PEM representation of the replica's x509
	   certificate.
	*/
	ClientCertificate string `json:"clientCertificate,omitempty" yaml:"clientCertificate,omitempty"`

	/*
	   PEM representation of the replica's private key. The
	   corresponding public key in encoded in the `client_certificate`.
	*/
	ClientKey string `json:"clientKey,omitempty" yaml:"clientKey,omitempty"`

	/*
	   The number of seconds
	   between connect retries. MySQL's default is 60 seconds.
	*/
	ConnectRetryInterval int `json:"connectRetryInterval,omitempty" yaml:"connectRetryInterval,omitempty"`
}
