package types

type Blockchainnodeengine_BlockchainNodesEthereumDetails struct {
	/*
	   The Ethereum environment being accessed.
	   Possible values are: `MAINNET`, `TESTNET_GOERLI_PRATER`, `TESTNET_SEPOLIA`.
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   The type of Ethereum node.
	   Possible values are: `LIGHT`, `FULL`, `ARCHIVE`.
	*/
	NodeType string `json:"nodeType,omitempty" yaml:"nodeType,omitempty"`

	/*
	   (Output)
	   User-provided key-value pairs
	   Structure is documented below.
	*/
	AdditionalEndpoints []Blockchainnodeengine_BlockchainNodesEthereumDetailsAdditionalEndpoint `json:"additionalEndpoints,omitempty" yaml:"additionalEndpoints,omitempty"`

	// Enables JSON-RPC access to functions in the admin namespace. Defaults to false.
	ApiEnableAdmin bool `json:"apiEnableAdmin,omitempty" yaml:"apiEnableAdmin,omitempty"`

	// Enables JSON-RPC access to functions in the debug namespace. Defaults to false.
	ApiEnableDebug bool `json:"apiEnableDebug,omitempty" yaml:"apiEnableDebug,omitempty"`

	/*
	   The consensus client
	   Possible values are: `CONSENSUS_CLIENT_UNSPECIFIED`, `LIGHTHOUSE`.
	*/
	ConsensusClient string `json:"consensusClient,omitempty" yaml:"consensusClient,omitempty"`

	/*
	   The execution client
	   Possible values are: `EXECUTION_CLIENT_UNSPECIFIED`, `GETH`, `ERIGON`.
	*/
	ExecutionClient string `json:"executionClient,omitempty" yaml:"executionClient,omitempty"`

	/*
	   User-provided key-value pairs
	   Structure is documented below.
	*/
	GethDetails Blockchainnodeengine_BlockchainNodesEthereumDetailsGethDetails `json:"gethDetails,omitempty" yaml:"gethDetails,omitempty"`

	/*
	   Configuration for validator-related parameters on the beacon client, and for any managed validator client.
	   Structure is documented below.
	*/
	ValidatorConfig Blockchainnodeengine_BlockchainNodesEthereumDetailsValidatorConfig `json:"validatorConfig,omitempty" yaml:"validatorConfig,omitempty"`
}
