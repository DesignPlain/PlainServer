package types

type Mskconnect_ConnectorKafkaClusterClientAuthentication struct {
	// The type of client authentication used to connect to the Apache Kafka cluster. Valid values: `IAM`, `NONE`. A value of `NONE` means that no client authentication is used. The default value is `NONE`.
	AuthenticationType string `json:"authenticationType,omitempty" yaml:"authenticationType,omitempty"`
}
