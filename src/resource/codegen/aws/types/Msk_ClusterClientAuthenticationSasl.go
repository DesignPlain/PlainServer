package types

type Msk_ClusterClientAuthenticationSasl struct {
	//
	Iam bool `json:"iam,omitempty" yaml:"iam,omitempty"`

	//
	Scram bool `json:"scram,omitempty" yaml:"scram,omitempty"`
}
