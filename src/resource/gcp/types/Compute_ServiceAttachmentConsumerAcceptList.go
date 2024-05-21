package types

type Compute_ServiceAttachmentConsumerAcceptList struct {
	// A project that is allowed to connect to this service attachment.
	ProjectIdOrNum string `json:"projectIdOrNum,omitempty" yaml:"projectIdOrNum,omitempty"`

	/*
	   The number of consumer forwarding rules the consumer project can
	   create.
	*/
	ConnectionLimit int `json:"connectionLimit,omitempty" yaml:"connectionLimit,omitempty"`
}
