package s3control

type AccessGrantsInstanceResourcePolicy struct {
	//
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// The policy document.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`
}
