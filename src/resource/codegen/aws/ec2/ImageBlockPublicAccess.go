package ec2

type ImageBlockPublicAccess struct {
	// The state of block public access for AMIs at the account level in the configured AWS Region. Valid values: `unblocked` and `block-new-sharing`.
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
