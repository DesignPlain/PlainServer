package opensearch

type InboundConnectionAccepter struct {
	// Specifies the ID of the connection to accept.
	ConnectionId string `json:"connectionId,omitempty" yaml:"connectionId,omitempty"`
}
