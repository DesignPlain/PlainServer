package types

type Msk_getClusterBrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthentication struct {
	//
	Sasls []Msk_getClusterBrokerNodeGroupInfoConnectivityInfoVpcConnectivityClientAuthenticationSasl `json:"sasls,omitempty" yaml:"sasls,omitempty"`

	//
	Tls bool `json:"tls,omitempty" yaml:"tls,omitempty"`
}
