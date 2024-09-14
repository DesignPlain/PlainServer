package types

type Mskconnect_ConnectorKafkaClusterEncryptionInTransit struct {
	// The type of encryption in transit to the Apache Kafka cluster. Valid values: `PLAINTEXT`, `TLS`. The default values is `PLAINTEXT`.
	EncryptionType string `json:"encryptionType,omitempty" yaml:"encryptionType,omitempty"`
}
