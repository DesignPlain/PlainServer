package types

type Beyondcorp_getAppGatewayAllocatedConnection struct {
	// The PSC uri of an allocated connection.
	PscUri string `json:"pscUri,omitempty" yaml:"pscUri,omitempty"`

	// The ingress port of an allocated connection.
	IngressPort int `json:"ingressPort,omitempty" yaml:"ingressPort,omitempty"`
}
