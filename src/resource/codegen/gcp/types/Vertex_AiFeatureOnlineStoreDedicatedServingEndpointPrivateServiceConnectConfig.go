package types

type Vertex_AiFeatureOnlineStoreDedicatedServingEndpointPrivateServiceConnectConfig struct {
	// If set to true, customers will use private service connection to send request. Otherwise, the connection will set to public endpoint.
	EnablePrivateServiceConnect bool `json:"enablePrivateServiceConnect,omitempty" yaml:"enablePrivateServiceConnect,omitempty"`

	// A list of Projects from which the forwarding rule will target the service attachment.
	ProjectAllowlists []string `json:"projectAllowlists,omitempty" yaml:"projectAllowlists,omitempty"`
}
