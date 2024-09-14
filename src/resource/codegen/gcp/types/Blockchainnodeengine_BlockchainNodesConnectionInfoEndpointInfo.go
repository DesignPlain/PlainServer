package types

type Blockchainnodeengine_BlockchainNodesConnectionInfoEndpointInfo struct {
	/*
	   (Output)
	   The assigned URL for the node JSON-RPC API endpoint.
	*/
	JsonRpcApiEndpoint string `json:"jsonRpcApiEndpoint,omitempty" yaml:"jsonRpcApiEndpoint,omitempty"`

	/*
	   (Output)
	   The assigned URL for the node WebSockets API endpoint.
	*/
	WebsocketsApiEndpoint string `json:"websocketsApiEndpoint,omitempty" yaml:"websocketsApiEndpoint,omitempty"`
}
