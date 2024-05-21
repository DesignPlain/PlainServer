package types

type Vertex_AiIndexEndpointPrivateServiceConnectConfig struct {
	// If set to true, the IndexEndpoint is created without private service access.
	EnablePrivateServiceConnect bool `json:"enablePrivateServiceConnect,omitempty" yaml:"enablePrivateServiceConnect,omitempty"`

	// A list of Projects from which the forwarding rule will target the service attachment.
	ProjectAllowlists []string `json:"projectAllowlists,omitempty" yaml:"projectAllowlists,omitempty"`
}
