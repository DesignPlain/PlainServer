package types

type Networkmanager_ConnectAttachmentOptions struct {
	// The protocol used for the attachment connection. Possible values are `GRE` and `NO_ENCAP`.
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`
}
