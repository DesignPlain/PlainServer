package types

type Opensearch_OutboundConnectionConnectionProperties struct {
	// The endpoint of the remote domain, is only set when `connection_mode` is `VPC_ENDPOINT` and `accept_connection` is `TRUE`.
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`

	// Configuration block for cross cluster search.
	CrossClusterSearch Opensearch_OutboundConnectionConnectionPropertiesCrossClusterSearch `json:"crossClusterSearch,omitempty" yaml:"crossClusterSearch,omitempty"`
}
