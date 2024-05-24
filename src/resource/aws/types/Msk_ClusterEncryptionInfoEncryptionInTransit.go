package types

type Msk_ClusterEncryptionInfoEncryptionInTransit struct {
	// Encryption setting for data in transit between clients and brokers. Valid values: `TLS`, `TLS_PLAINTEXT`, and `PLAINTEXT`. Default value is `TLS`.
	ClientBroker string `json:"clientBroker,omitempty" yaml:"clientBroker,omitempty"`

	// Whether data communication among broker nodes is encrypted. Default value: `true`.
	InCluster bool `json:"inCluster,omitempty" yaml:"inCluster,omitempty"`
}
