package types

type Blockchainnodeengine_BlockchainNodesEthereumDetailsAdditionalEndpoint struct {
	/*
	   (Output)
	   The assigned URL for the node's Beacon API endpoint.
	*/
	BeaconApiEndpoint string `json:"beaconApiEndpoint,omitempty" yaml:"beaconApiEndpoint,omitempty"`

	/*
	   (Output)
	   The assigned URL for the node's Beacon Prometheus metrics endpoint.
	*/
	BeaconPrometheusMetricsApiEndpoint string `json:"beaconPrometheusMetricsApiEndpoint,omitempty" yaml:"beaconPrometheusMetricsApiEndpoint,omitempty"`

	/*
	   (Output)
	   The assigned URL for the node's execution client's Prometheus metrics endpoint.
	*/
	ExecutionClientPrometheusMetricsApiEndpoint string `json:"executionClientPrometheusMetricsApiEndpoint,omitempty" yaml:"executionClientPrometheusMetricsApiEndpoint,omitempty"`
}
