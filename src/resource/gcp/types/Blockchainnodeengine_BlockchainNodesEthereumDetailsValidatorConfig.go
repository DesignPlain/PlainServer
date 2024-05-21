package types

type Blockchainnodeengine_BlockchainNodesEthereumDetailsValidatorConfig struct {
	// URLs for MEV-relay services to use for block building. When set, a managed MEV-boost service is configured on the beacon client.
	MevRelayUrls []string `json:"mevRelayUrls,omitempty" yaml:"mevRelayUrls,omitempty"`
}
