package blockchainnodeengine

import types "libds/gcp/types"

type BlockchainNodes struct {
	/*
	   ID of the requesting object.


	   - - -
	*/
	BlockchainNodeId string `json:"blockchainNodeId,omitempty" yaml:"blockchainNodeId,omitempty"`

	/*
	   User-provided key-value pairs
	   Possible values are: `ETHEREUM`.
	*/
	BlockchainType string `json:"blockchainType,omitempty" yaml:"blockchainType,omitempty"`

	/*
	   User-provided key-value pairs
	   Structure is documented below.
	*/
	EthereumDetails types.Blockchainnodeengine_BlockchainNodesEthereumDetails `json:"ethereumDetails,omitempty" yaml:"ethereumDetails,omitempty"`

	/*
	   User-provided key-value pairs

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// Location of Blockchain Node being created.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
