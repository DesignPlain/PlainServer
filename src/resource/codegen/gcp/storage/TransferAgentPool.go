package storage

import types "libds/gcp/types"

type TransferAgentPool struct {
	// Specifies the client-specified AgentPool description.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   The ID of the agent pool to create.
	   The agentPoolId must meet the following requirements:
	   - Length of 128 characters or less.
	   - Not start with the string goog.
	   - Start with a lowercase ASCII character, followed by:
	   - Zero or more: lowercase Latin alphabet characters, numerals, hyphens (-), periods (.), underscores (_), or tildes (~).
	   - One or more numerals or lowercase ASCII characters.
	   As expressed by the regular expression: ^(?!goog)a-z?$.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Specifies the bandwidth limit details. If this field is unspecified, the default value is set as 'No Limit'.
	   Structure is documented below.
	*/
	BandwidthLimit types.Storage_TransferAgentPoolBandwidthLimit `json:"bandwidthLimit,omitempty" yaml:"bandwidthLimit,omitempty"`
}
