package types

type Msk_ServerlessClusterClientAuthenticationSasl struct {
	// Details for client authentication using IAM. See below.
	Iam Msk_ServerlessClusterClientAuthenticationSaslIam `json:"iam,omitempty" yaml:"iam,omitempty"`
}
