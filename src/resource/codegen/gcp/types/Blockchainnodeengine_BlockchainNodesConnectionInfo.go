package types

type Blockchainnodeengine_BlockchainNodesConnectionInfo struct {
	/*
	   (Output)
	   The endpoint information through which to interact with a blockchain node.
	   Structure is documented below.
	*/
	EndpointInfos []Blockchainnodeengine_BlockchainNodesConnectionInfoEndpointInfo `json:"endpointInfos,omitempty" yaml:"endpointInfos,omitempty"`

	/*
	   (Output)
	   A service attachment that exposes a node, and has the following format: projects/{project}/regions/{region}/serviceAttachments/{service_attachment_name}
	*/
	ServiceAttachment string `json:"serviceAttachment,omitempty" yaml:"serviceAttachment,omitempty"`
}
