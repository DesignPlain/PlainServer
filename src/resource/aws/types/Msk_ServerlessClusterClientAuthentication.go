package types

type Msk_ServerlessClusterClientAuthentication struct {
	// Details for client authentication using SASL. See below.
	Sasl Msk_ServerlessClusterClientAuthenticationSasl `json:"sasl,omitempty" yaml:"sasl,omitempty"`
}
