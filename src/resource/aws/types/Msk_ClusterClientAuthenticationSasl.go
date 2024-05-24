package types

type Msk_ClusterClientAuthenticationSasl struct {
	// Enables SCRAM client authentication via AWS Secrets Manager. Defaults to `false`.
	Scram bool `json:"scram,omitempty" yaml:"scram,omitempty"`

	// Enables IAM client authentication. Defaults to `false`.
	Iam bool `json:"iam,omitempty" yaml:"iam,omitempty"`
}
