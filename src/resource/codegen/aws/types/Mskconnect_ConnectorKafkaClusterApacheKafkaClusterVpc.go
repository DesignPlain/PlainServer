package types

type Mskconnect_ConnectorKafkaClusterApacheKafkaClusterVpc struct {
	// The security groups for the connector.
	SecurityGroups []string `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`

	// The subnets for the connector.
	Subnets []string `json:"subnets,omitempty" yaml:"subnets,omitempty"`
}
