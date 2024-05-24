package directconnect

type ConnectionAssociation struct {
	// The ID of the connection.
	ConnectionId string `json:"connectionId,omitempty" yaml:"connectionId,omitempty"`

	// The ID of the LAG with which to associate the connection.
	LagId string `json:"lagId,omitempty" yaml:"lagId,omitempty"`
}
