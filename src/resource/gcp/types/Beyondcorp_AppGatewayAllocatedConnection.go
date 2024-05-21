package types

type Beyondcorp_AppGatewayAllocatedConnection struct {
	// The ingress port of an allocated connection.
	IngressPort int `json:"ingressPort,omitempty" yaml:"ingressPort,omitempty"`

	// The PSC uri of an allocated connection.
	PscUri string `json:"pscUri,omitempty" yaml:"pscUri,omitempty"`
}
