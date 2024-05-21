package types

type Datastream_ConnectionProfilePrivateConnectivity struct {
	// A reference to a private connection resource. Format: `projects/{project}/locations/{location}/privateConnections/{name}`
	PrivateConnection string `json:"privateConnection,omitempty" yaml:"privateConnection,omitempty"`
}
