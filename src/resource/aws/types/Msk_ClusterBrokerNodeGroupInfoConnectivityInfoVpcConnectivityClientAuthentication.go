package types

type Msk_ClusterBrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthentication struct {
	// Configuration block for specifying SASL client authentication. See below.
	Sasl Msk_ClusterBrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationSasl `json:"sasl,omitempty" yaml:"sasl,omitempty"`

	// Configuration block for specifying TLS client authentication. See below.
	Tls bool `json:"tls,omitempty" yaml:"tls,omitempty"`
}
