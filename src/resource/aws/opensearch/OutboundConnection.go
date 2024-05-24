package opensearch

import types "DesignSphere_Server/src/resource/aws/types"

type OutboundConnection struct {
	// Specifies the connection alias that will be used by the customer for this connection.
	ConnectionAlias string `json:"connectionAlias,omitempty" yaml:"connectionAlias,omitempty"`

	// Specifies the connection mode. Accepted values are `DIRECT` or `VPC_ENDPOINT`.
	ConnectionMode string `json:"connectionMode,omitempty" yaml:"connectionMode,omitempty"`

	// Configuration block for the outbound connection.
	ConnectionProperties types.Opensearch_OutboundConnectionConnectionProperties `json:"connectionProperties,omitempty" yaml:"connectionProperties,omitempty"`

	// Configuration block for the local Opensearch domain.
	LocalDomainInfo types.Opensearch_OutboundConnectionLocalDomainInfo `json:"localDomainInfo,omitempty" yaml:"localDomainInfo,omitempty"`

	// Configuration block for the remote Opensearch domain.
	RemoteDomainInfo types.Opensearch_OutboundConnectionRemoteDomainInfo `json:"remoteDomainInfo,omitempty" yaml:"remoteDomainInfo,omitempty"`

	// Accepts the connection.
	AcceptConnection bool `json:"acceptConnection,omitempty" yaml:"acceptConnection,omitempty"`
}
